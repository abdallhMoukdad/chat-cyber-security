// Code generated by ent, DO NOT EDIT.

package student

import (
	"awesomeProject1/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldName, v))
}

// EnrollmentDate applies equality check predicate on the "enrollment_date" field. It's identical to EnrollmentDateEQ.
func EnrollmentDate(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEnrollmentDate, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPassword, v))
}

// NationalNumber applies equality check predicate on the "national_number" field. It's identical to NationalNumberEQ.
func NationalNumber(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldNationalNumber, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPhoneNumber, v))
}

// HomeLocation applies equality check predicate on the "home_location" field. It's identical to HomeLocationEQ.
func HomeLocation(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldHomeLocation, v))
}

// EnycrptionKey applies equality check predicate on the "enycrption_key" field. It's identical to EnycrptionKeyEQ.
func EnycrptionKey(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEnycrptionKey, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldName, v))
}

// EnrollmentDateEQ applies the EQ predicate on the "enrollment_date" field.
func EnrollmentDateEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEnrollmentDate, v))
}

// EnrollmentDateNEQ applies the NEQ predicate on the "enrollment_date" field.
func EnrollmentDateNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldEnrollmentDate, v))
}

// EnrollmentDateIn applies the In predicate on the "enrollment_date" field.
func EnrollmentDateIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldEnrollmentDate, vs...))
}

// EnrollmentDateNotIn applies the NotIn predicate on the "enrollment_date" field.
func EnrollmentDateNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldEnrollmentDate, vs...))
}

// EnrollmentDateGT applies the GT predicate on the "enrollment_date" field.
func EnrollmentDateGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldEnrollmentDate, v))
}

// EnrollmentDateGTE applies the GTE predicate on the "enrollment_date" field.
func EnrollmentDateGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldEnrollmentDate, v))
}

// EnrollmentDateLT applies the LT predicate on the "enrollment_date" field.
func EnrollmentDateLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldEnrollmentDate, v))
}

// EnrollmentDateLTE applies the LTE predicate on the "enrollment_date" field.
func EnrollmentDateLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldEnrollmentDate, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldPassword, v))
}

// NationalNumberEQ applies the EQ predicate on the "national_number" field.
func NationalNumberEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldNationalNumber, v))
}

// NationalNumberNEQ applies the NEQ predicate on the "national_number" field.
func NationalNumberNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldNationalNumber, v))
}

// NationalNumberIn applies the In predicate on the "national_number" field.
func NationalNumberIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldNationalNumber, vs...))
}

// NationalNumberNotIn applies the NotIn predicate on the "national_number" field.
func NationalNumberNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldNationalNumber, vs...))
}

// NationalNumberGT applies the GT predicate on the "national_number" field.
func NationalNumberGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldNationalNumber, v))
}

// NationalNumberGTE applies the GTE predicate on the "national_number" field.
func NationalNumberGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldNationalNumber, v))
}

// NationalNumberLT applies the LT predicate on the "national_number" field.
func NationalNumberLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldNationalNumber, v))
}

// NationalNumberLTE applies the LTE predicate on the "national_number" field.
func NationalNumberLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldNationalNumber, v))
}

// NationalNumberContains applies the Contains predicate on the "national_number" field.
func NationalNumberContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldNationalNumber, v))
}

// NationalNumberHasPrefix applies the HasPrefix predicate on the "national_number" field.
func NationalNumberHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldNationalNumber, v))
}

// NationalNumberHasSuffix applies the HasSuffix predicate on the "national_number" field.
func NationalNumberHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldNationalNumber, v))
}

// NationalNumberEqualFold applies the EqualFold predicate on the "national_number" field.
func NationalNumberEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldNationalNumber, v))
}

// NationalNumberContainsFold applies the ContainsFold predicate on the "national_number" field.
func NationalNumberContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldNationalNumber, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// HomeLocationEQ applies the EQ predicate on the "home_location" field.
func HomeLocationEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldHomeLocation, v))
}

// HomeLocationNEQ applies the NEQ predicate on the "home_location" field.
func HomeLocationNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldHomeLocation, v))
}

// HomeLocationIn applies the In predicate on the "home_location" field.
func HomeLocationIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldHomeLocation, vs...))
}

// HomeLocationNotIn applies the NotIn predicate on the "home_location" field.
func HomeLocationNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldHomeLocation, vs...))
}

// HomeLocationGT applies the GT predicate on the "home_location" field.
func HomeLocationGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldHomeLocation, v))
}

// HomeLocationGTE applies the GTE predicate on the "home_location" field.
func HomeLocationGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldHomeLocation, v))
}

// HomeLocationLT applies the LT predicate on the "home_location" field.
func HomeLocationLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldHomeLocation, v))
}

// HomeLocationLTE applies the LTE predicate on the "home_location" field.
func HomeLocationLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldHomeLocation, v))
}

// HomeLocationContains applies the Contains predicate on the "home_location" field.
func HomeLocationContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldHomeLocation, v))
}

// HomeLocationHasPrefix applies the HasPrefix predicate on the "home_location" field.
func HomeLocationHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldHomeLocation, v))
}

// HomeLocationHasSuffix applies the HasSuffix predicate on the "home_location" field.
func HomeLocationHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldHomeLocation, v))
}

// HomeLocationEqualFold applies the EqualFold predicate on the "home_location" field.
func HomeLocationEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldHomeLocation, v))
}

// HomeLocationContainsFold applies the ContainsFold predicate on the "home_location" field.
func HomeLocationContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldHomeLocation, v))
}

// EnycrptionKeyEQ applies the EQ predicate on the "enycrption_key" field.
func EnycrptionKeyEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEnycrptionKey, v))
}

// EnycrptionKeyNEQ applies the NEQ predicate on the "enycrption_key" field.
func EnycrptionKeyNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldEnycrptionKey, v))
}

// EnycrptionKeyIn applies the In predicate on the "enycrption_key" field.
func EnycrptionKeyIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldEnycrptionKey, vs...))
}

// EnycrptionKeyNotIn applies the NotIn predicate on the "enycrption_key" field.
func EnycrptionKeyNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldEnycrptionKey, vs...))
}

// EnycrptionKeyGT applies the GT predicate on the "enycrption_key" field.
func EnycrptionKeyGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldEnycrptionKey, v))
}

// EnycrptionKeyGTE applies the GTE predicate on the "enycrption_key" field.
func EnycrptionKeyGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldEnycrptionKey, v))
}

// EnycrptionKeyLT applies the LT predicate on the "enycrption_key" field.
func EnycrptionKeyLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldEnycrptionKey, v))
}

// EnycrptionKeyLTE applies the LTE predicate on the "enycrption_key" field.
func EnycrptionKeyLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldEnycrptionKey, v))
}

// EnycrptionKeyContains applies the Contains predicate on the "enycrption_key" field.
func EnycrptionKeyContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldEnycrptionKey, v))
}

// EnycrptionKeyHasPrefix applies the HasPrefix predicate on the "enycrption_key" field.
func EnycrptionKeyHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldEnycrptionKey, v))
}

// EnycrptionKeyHasSuffix applies the HasSuffix predicate on the "enycrption_key" field.
func EnycrptionKeyHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldEnycrptionKey, v))
}

// EnycrptionKeyEqualFold applies the EqualFold predicate on the "enycrption_key" field.
func EnycrptionKeyEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldEnycrptionKey, v))
}

// EnycrptionKeyContainsFold applies the ContainsFold predicate on the "enycrption_key" field.
func EnycrptionKeyContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldEnycrptionKey, v))
}

// HasChats applies the HasEdge predicate on the "chats" edge.
func HasChats() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ChatsTable, ChatsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChatsWith applies the HasEdge predicate on the "chats" edge with a given conditions (other predicates).
func HasChatsWith(preds ...predicate.Chat) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := newChatsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSentMessages applies the HasEdge predicate on the "sent_messages" edge.
func HasSentMessages() predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SentMessagesTable, SentMessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSentMessagesWith applies the HasEdge predicate on the "sent_messages" edge with a given conditions (other predicates).
func HasSentMessagesWith(preds ...predicate.Message) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		step := newSentMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(sql.NotPredicates(p))
}
