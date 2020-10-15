// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/B6113759/app/ent/coolroom"
	"github.com/B6113759/app/ent/coolroomtype"
	"github.com/B6113759/app/ent/patient"
	"github.com/B6113759/app/ent/relative"
	"github.com/B6113759/app/ent/schema"
	"github.com/B6113759/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	coolroomFields := schema.Coolroom{}.Fields()
	_ = coolroomFields
	// coolroomDescCoolroomName is the schema descriptor for coolroom_name field.
	coolroomDescCoolroomName := coolroomFields[0].Descriptor()
	// coolroom.CoolroomNameValidator is a validator for the "coolroom_name" field. It is called by the builders before save.
	coolroom.CoolroomNameValidator = coolroomDescCoolroomName.Validators[0].(func(string) error)
	coolroomtypeFields := schema.CoolroomType{}.Fields()
	_ = coolroomtypeFields
	// coolroomtypeDescCoolroomtypeName is the schema descriptor for coolroomtype_name field.
	coolroomtypeDescCoolroomtypeName := coolroomtypeFields[0].Descriptor()
	// coolroomtype.CoolroomtypeNameValidator is a validator for the "coolroomtype_name" field. It is called by the builders before save.
	coolroomtype.CoolroomtypeNameValidator = coolroomtypeDescCoolroomtypeName.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for patient_name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "patient_name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	// patientDescPatientAge is the schema descriptor for patient_age field.
	patientDescPatientAge := patientFields[1].Descriptor()
	// patient.PatientAgeValidator is a validator for the "patient_age" field. It is called by the builders before save.
	patient.PatientAgeValidator = patientDescPatientAge.Validators[0].(func(int) error)
	relativeFields := schema.Relative{}.Fields()
	_ = relativeFields
	// relativeDescRelativeName is the schema descriptor for relative_name field.
	relativeDescRelativeName := relativeFields[0].Descriptor()
	// relative.RelativeNameValidator is a validator for the "relative_name" field. It is called by the builders before save.
	relative.RelativeNameValidator = relativeDescRelativeName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[0].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
	// userDescUseremail is the schema descriptor for useremail field.
	userDescUseremail := userFields[1].Descriptor()
	// user.UseremailValidator is a validator for the "useremail" field. It is called by the builders before save.
	user.UseremailValidator = userDescUseremail.Validators[0].(func(string) error)
}
