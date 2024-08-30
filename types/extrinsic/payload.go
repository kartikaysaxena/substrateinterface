package extrinsic

import (
	"errors"
	"fmt"

	"github.com/kartikaysaxena/substrateinterface/scale"
	"github.com/kartikaysaxena/substrateinterface/signature"
	"github.com/kartikaysaxena/substrateinterface/types"
	"github.com/kartikaysaxena/substrateinterface/types/codec"
	"github.com/kartikaysaxena/substrateinterface/types/extrinsic/extensions"
)

// SignedField represents a field used in the Payload.
type SignedField struct {
	// Name of the field, this is omitted when the extrinsic is signed.
	Name SignedFieldName

	// Value of the field used when the extrinsic is signed.
	Value any

	// Mutated is used to keep track of changes done to the field.
	Mutated bool
}

// Payload holds the encoded types.Call and the fields that are used for generating
// the Extrinsic payload and signature.
//
// Notes - the ordering of the SignedFields and SignedExtraFields is the order in which they are provided in
// the metadata.
type Payload struct {
	EncodedCall       types.BytesBare
	SignedFields      []*SignedField
	SignedExtraFields []*SignedField
}

// Encode the call bytes and the signed fields in the order that is provided during creation
// from the metadata.
//
// The function also performs an extra check to ensure that all required fields were mutated.
func (p *Payload) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(p.EncodedCall); err != nil {
		return fmt.Errorf("unable to encode method: %w", err)
	}

	for _, signedField := range p.SignedFields {
		if !signedField.Mutated {
			return fmt.Errorf("signed field '%s' was not mutated", signedField.Name)
		}

		if err := encoder.Encode(signedField.Value); err != nil {
			return fmt.Errorf("unable to encode signed field: %w", err)
		}
	}

	for _, signedExtraField := range p.SignedExtraFields {
		if !signedExtraField.Mutated {
			return fmt.Errorf("signed extra field '%s' was not mutated", signedExtraField.Name)
		}

		if err := encoder.Encode(signedExtraField.Value); err != nil {
			return fmt.Errorf("unable to encode signed extra field: %w", err)
		}
	}

	return nil
}

// MutateSignedFields is mutating the payload's SignedFields and SignedExtraFields
// based on the provided SignedFieldValues.
func (p *Payload) MutateSignedFields(vals SignedFieldValues) error {
	if p == nil {
		return errors.New("payload is nil")
	}

	for _, signedField := range p.SignedFields {
		signedFieldVal, ok := vals[signedField.Name]

		if !ok {
			continue
		}

		signedField.Value = signedFieldVal
		signedField.Mutated = true
	}

	for _, signedExtraField := range p.SignedExtraFields {
		signedFieldVal, ok := vals[signedExtraField.Name]

		if !ok {
			continue
		}

		signedExtraField.Value = signedFieldVal
		signedExtraField.Mutated = true
	}

	return nil
}

// Sign encodes the payload and then signs the encoded bytes using the provided signer.
func (p *Payload) Sign(signer signature.KeyringPair) (types.SignatureHash, error) {
	b, err := codec.Encode(p)
	if err != nil {
		return types.SignatureHash{}, err
	}

	sig, err := signature.Sign(b, signer.URI)
	return types.NewSignature(sig), err
}

// SignedFieldName is the type used for representing a field name.
type SignedFieldName string

const (
	EraSignedField                   SignedFieldName = "era"
	BlockHashSignedField             SignedFieldName = "block_hash"
	NonceSignedField                 SignedFieldName = "nonce"
	TipSignedField                   SignedFieldName = "tip"
	AssetIDSignedField               SignedFieldName = "asset_id"
	CheckMetadataHashModeSignedField SignedFieldName = "check_metadata_hash_mode"
	CheckMetadataHashSignedField     SignedFieldName = "check_metadata_hash"
	SpecVersionSignedField           SignedFieldName = "spec_version"
	TransactionVersionSignedField    SignedFieldName = "transaction_version"
	GenesisHashSignedField           SignedFieldName = "genesis_hash"
)

// PayloadMutatorFn is the type used for mutating the Payload during creation.
type PayloadMutatorFn func(payload *Payload)

// PayloadMutatorFns is a map that holds a PayloadMutatorFn for each supported signed extension.
var PayloadMutatorFns = map[extensions.SignedExtensionName]PayloadMutatorFn{
	extensions.CheckMortalitySignedExtension: func(payload *Payload) {
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  EraSignedField,
			Value: &types.ExtrinsicEra{},
		})
		payload.SignedExtraFields = append(payload.SignedExtraFields, &SignedField{
			Name:  BlockHashSignedField,
			Value: &types.Hash{},
		})
	},
	extensions.CheckEraSignedExtension: func(payload *Payload) {
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  EraSignedField,
			Value: &types.ExtrinsicEra{},
		})
		payload.SignedExtraFields = append(payload.SignedExtraFields, &SignedField{
			Name:  BlockHashSignedField,
			Value: &types.Hash{},
		})
	},
	extensions.CheckNonceSignedExtension: func(payload *Payload) {
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  NonceSignedField,
			Value: types.U32(0),
		})
	},
	extensions.ChargeTransactionPaymentSignedExtension: func(payload *Payload) {
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  TipSignedField,
			Value: types.NewUCompactFromUInt(0),
		})
	},
	extensions.ChargeAssetTxPaymentSignedExtension: func(payload *Payload) {
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  TipSignedField,
			Value: types.NewUCompactFromUInt(0),
		})
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  AssetIDSignedField,
			Value: types.NewEmptyOption[types.AssetID](),
		})
	},
	extensions.CheckMetadataHashSignedExtension: func(payload *Payload) {
		payload.SignedFields = append(payload.SignedFields, &SignedField{
			Name:  CheckMetadataHashModeSignedField,
			Value: extensions.CheckMetadataModeDisabled,
		})
		payload.SignedExtraFields = append(payload.SignedExtraFields, &SignedField{
			Name:  CheckMetadataHashSignedField,
			Value: &extensions.CheckMetadataHash{},
		})
	},
	extensions.CheckSpecVersionSignedExtension: func(payload *Payload) {
		payload.SignedExtraFields = append(payload.SignedExtraFields, &SignedField{
			Name:  SpecVersionSignedField,
			Value: types.U32(0),
		})
	},
	extensions.CheckTxVersionSignedExtension: func(payload *Payload) {
		payload.SignedExtraFields = append(payload.SignedExtraFields, &SignedField{
			Name:  TransactionVersionSignedField,
			Value: types.U32(0),
		})
	},
	extensions.CheckGenesisSignedExtension: func(payload *Payload) {
		payload.SignedExtraFields = append(payload.SignedExtraFields, &SignedField{
			Name:  GenesisHashSignedField,
			Value: &types.Hash{},
		})
	},
	// There's nothing that we can add in the payload or signature in the following cases, however, these are added to
	// ensure that the extension is acknowledged and that the mutator check is passing.
	extensions.CheckNonZeroSenderSignedExtension:          func(payload *Payload) {},
	extensions.CheckWeightSignedExtension:                 func(payload *Payload) {},
	extensions.PreBalanceTransferExtensionSignedExtension: func(payload *Payload) {},
	extensions.StorageWeightReclaimSignedExtension:        func(payload *Payload) {},
	extensions.PrevalidateAttestsSignedExtension:          func(payload *Payload) {},
	extensions.CheckNetworkMembershipSignedExtension:      func(payload *Payload) {},
}

// createPayload iterates over all signed extensions provided in the metadata and
// attempts to load and use a PayloadMutatorFn for each one.
//
// If a PayloadMutatorFn is not found for a specific signed extension, it means that it is not currently supported.
func createPayload(meta *types.Metadata, encodedCall []byte) (*Payload, error) {
	payload := &Payload{
		EncodedCall: encodedCall,
	}

	for _, signedExtension := range meta.AsMetadataV14.Extrinsic.SignedExtensions {
		signedExtensionType, ok := meta.AsMetadataV14.EfficientLookup[signedExtension.Type.Int64()]

		if !ok {
			return nil, fmt.Errorf("signed extension type '%d' is not defined", signedExtension.Type.Int64())
		}

		signedExtensionName := extensions.SignedExtensionName(signedExtensionType.Path[len(signedExtensionType.Path)-1])

		payloadMutatorFn, ok := PayloadMutatorFns[signedExtensionName]

		if !ok {
			return nil, fmt.Errorf("signed extension '%s' is not supported", signedExtensionName)
		}

		payloadMutatorFn(payload)
	}

	return payload, nil
}
