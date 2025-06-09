package openapi

type AuthenticationSubscription struct {
	AuthenticationMethod               AuthType        `json:"authenticationMethod" bson:"authenticationMethod"`
	PermanentKey                       *PermanentKey   `json:"permanentKey" bson:"permanentKey"`
	SequenceNumber                     string          `json:"sequenceNumber" bson:"sequenceNumber"`
	AuthenticationManagementField      string          `json:"authenticationManagementField,omitempty" bson:"authenticationManagementField"`
	VectorAlgorithm                    VectorAlgorithm `json:"vectorAlgorithm,omitempty" bson:"vectorAlgorithm"`
	Milenage                           *Milenage       `json:"milenage,omitempty" bson:"milenage"`
	Tuak                               *Tuak           `json:"tuak,omitempty" bson:"tuak"`
	Opc                                *Opc            `json:"opc,omitempty" bson:"opc"`
	Topc                               *Topc           `json:"topc,omitempty" bson:"topc"`
	SharedAuthenticationSubscriptionId *SharedData     `json:"sharedAuthenticationSubscriptionId,omitempty" bson:"sharedAuthenticationSubscriptionId"`
}

// type PermanentKey struct {
// 	PermanentKeyValue   string `json:"permanentKeyValue" bson:"permanentKeyValue"`
// 	EncryptionKey       int32  `json:"encryptionKey" bson:"encryptionKey"`
// 	EncryptionAlgorithm int32  `json:"encryptionAlgorithm" bson:"encryptionAlgorithm"`
// }

// type VectorAlgorithm string

// const (
// 	VectorAlgorithm_MILENAGE VectorAlgorithm = "MILENAGE"
// 	VectorAlgorithm_TUAK     VectorAlgorithm = "TUAK"
// )

// type Milenage struct {
// 	Op        *Op        `json:"op,omitempty" bson:"op"`
// 	Rotations *Rotations `json:"rotations,omitempty" bson:"rotations"`
// 	Constants *Constants `json:"constants,omitempty" bson:"constants"`
// }

// type Op struct {
// 	OpValue             string `json:"opValue" bson:"opValue"`
// 	EncryptionKey       int32  `json:"encryptionKey" bson:"encryptionKey"`
// 	EncryptionAlgorithm int32  `json:"encryptionAlgorithm" bson:"encryptionAlgorithm"`
// }

// type Rotations struct {
// 	R1 string `json:"r1" bson:"r1"`
// 	R2 string `json:"r2" bson:"r2"`
// 	R3 string `json:"r3" bson:"r3"`
// 	R4 string `json:"r4" bson:"r4"`
// 	R5 string `json:"r5" bson:"r5"`
// }

// type Constants struct {
// 	C1 string `json:"c1" bson:"c1"`
// 	C2 string `json:"c2" bson:"c2"`
// 	C3 string `json:"c3" bson:"c3"`
// 	C4 string `json:"c4" bson:"c4"`
// 	C5 string `json:"c5" bson:"c5"`
// }

// type Tuak struct {
// 	Top              *Top  `json:"top,omitempty" bson:"top"`
// 	KeccakIterations int32 `json:"keccakIterations,omitempty" bson:"keccakIterations"`
// }
// type Top struct {
// 	TopValue            string `json:"topValue" bson:"topValue"`
// 	EncryptionKey       int32  `json:"encryptionKey" bson:"encryptionKey"`
// 	EncryptionAlgorithm int32  `json:"encryptionAlgorithm" bson:"encryptionAlgorithm"`
// }

// type Opc struct {
// 	OpcValue            string `json:"opcValue" bson:"opcValue"`
// 	EncryptionKey       int32  `json:"encryptionKey" bson:"encryptionKey"`
// 	EncryptionAlgorithm int32  `json:"encryptionAlgorithm" bson:"encryptionAlgorithm"`
// }

// type Topc struct {
// 	TopcValue           string `json:"topcValue" bson:"topcValue"`
// 	EncryptionKey       int32  `json:"encryptionKey" bson:"encryptionKey"`
// 	EncryptionAlgorithm int32  `json:"encryptionAlgorithm" bson:"encryptionAlgorithm"`
// }
