package validation_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

const (
	minNamespaceLength = 3
	minRelationLength  = 3

	maxCaveatName      = 128
	maxTenantLength    = 63
	maxNamespaceLength = 64
	maxObjectIDLength  = 128
	maxRelationLength  = 64
)

var namespaces = []struct {
	name  string
	valid bool
}{
	{"", false},
	{"...", false},
	{"foo", true},
	{"bar", true},
	{"foo1", true},
	{"bar1", true},
	{"ab", false},
	{"Foo1", false},
	{"foo_bar", true},
	{"foo_bar_", false},
	{"foo/bar", true},
	{"foo/b", false},
	{"Foo/bar", false},
	{"foo/bar/baz", false},
	{strings.Repeat("f", minNamespaceLength-1), false},
	{strings.Repeat("f", minNamespaceLength), true},
	{strings.Repeat("\u0394", minNamespaceLength), false},
	{strings.Repeat("\n", minNamespaceLength), false},
	{strings.Repeat("_", minNamespaceLength), false},
	{strings.Repeat("-", minNamespaceLength), false},
	{strings.Repeat("/", minNamespaceLength), false},
	{strings.Repeat("\\", minNamespaceLength), false},
	{strings.Repeat("f", maxNamespaceLength), true},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", maxTenantLength), strings.Repeat("f", maxNamespaceLength)), true},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", maxTenantLength+1), strings.Repeat("f", maxNamespaceLength)), false},
	{fmt.Sprintf("%s/%s", strings.Repeat("f", maxTenantLength), strings.Repeat("f", maxNamespaceLength+1)), false},
	{strings.Repeat("f", maxNamespaceLength+1), false},
}

var objectIDs = []struct {
	name  string
	valid bool
}{
	{"a", true},
	{"1", true},
	{"a-", true},
	{"A-A-A", true},
	{"123e4567-e89b-12d3-a456-426614174000", true},
	{strings.Repeat("f", 64), true},
	{strings.Repeat("f", maxObjectIDLength), true},
	{"", false},
	{"  ", false},
	{"-", false},
	{strings.Repeat("\u0394", 4), false},
	{strings.Repeat("f", maxObjectIDLength+1), false},
	{"a@b.com", false},
	{"test@example.com", false},
	{"test*", false},
	{"*test", false},
	{"t*t", false},
	{"test@example.*", false},
	{"authn|someauthnvalue", true},
	{"authn|a@b.com", false},
	{"authn|", true},
	{"authn|a-b-c-d-e", true},
	{"authn|a_b", true},
}

var subjectIDs = append([]struct {
	name  string
	valid bool
}{
	{"*", true},
}, objectIDs...)

var caveats = []struct {
	name    string
	context map[string]any
	valid   bool
}{
	{"test", map[string]any{}, true},
	{"", map[string]any{"a": "b"}, false},
	{strings.Repeat("f", maxCaveatName), map[string]any{"a": "b"}, true},
	{strings.Repeat("f", maxCaveatName+1), map[string]any{"a": "b"}, false},
	{"test", map[string]any{"a": "b"}, true},
	{"test", nil, true},
	{"test", generateMap(128), true},
}

type relationValidity int

const (
	alwaysInvalid relationValidity = iota
	alwaysValid
	validV0SubjectOnly
	validV1SubjectOnly
)

type relationEntry struct {
	name     string
	validity relationValidity
}

var knownGoodONR = &v0.ObjectAndRelation{
	Namespace: "user",
	ObjectId:  "testuser",
	Relation:  "member",
}

var knownGoodObjectRef = &v1.ObjectReference{
	ObjectType: "user",
	ObjectId:   "testuser",
}

var knownGoodSubjectRef = &v1.SubjectReference{
	Object:           knownGoodObjectRef,
	OptionalRelation: "member",
}

var relations = []relationEntry{
	{"...", validV0SubjectOnly},
	{"", validV1SubjectOnly},
	{"foo", alwaysValid},
	{"bar", alwaysValid},
	{"foo1", alwaysValid},
	{"bar1", alwaysValid},
	{"ab", alwaysInvalid},
	{"Foo1", alwaysInvalid},
	{"foo_bar", alwaysValid},
	{"foo_bar_", alwaysInvalid},
	{"foo/bar", alwaysInvalid},
	{"foo/b", alwaysInvalid},
	{"Foo/bar", alwaysInvalid},
	{"foo/bar/baz", alwaysInvalid},
	{strings.Repeat("f", minRelationLength-1), alwaysInvalid},
	{strings.Repeat("f", minRelationLength), alwaysValid},
	{strings.Repeat("\u0394", minRelationLength), alwaysInvalid},
	{strings.Repeat("\n", minRelationLength), alwaysInvalid},
	{strings.Repeat("_", minRelationLength), alwaysInvalid},
	{strings.Repeat("-", minRelationLength), alwaysInvalid},
	{strings.Repeat("/", minRelationLength), alwaysInvalid},
	{strings.Repeat("\\", minRelationLength), alwaysInvalid},
	{strings.Repeat("f", maxRelationLength), alwaysValid},
	{strings.Repeat("f", maxRelationLength+1), alwaysInvalid},
}

func TestV0CoreObjectValidity(t *testing.T) {
	for _, ns := range namespaces {
		for _, objectID := range objectIDs {
			for _, subjectID := range objectIDs {
				for _, relation := range relations {
					testName := fmt.Sprintf("%s:%s#%s@%s", ns.name, objectID.name, relation.name, subjectID.name)
					t.Run(testName, func(t *testing.T) {
						t.Parallel()

						require := require.New(t)

						v0ObjectValid := ns.valid && objectID.valid && (relation.validity == alwaysValid ||
							relation.validity == validV0SubjectOnly)
						v0SubjectValid := ns.valid && subjectID.valid && (relation.validity == alwaysValid ||
							relation.validity == validV0SubjectOnly)
						v0Valid := v0ObjectValid && v0SubjectValid

						onr := &v0.ObjectAndRelation{
							Namespace: ns.name,
							ObjectId:  objectID.name,
							Relation:  relation.name,
						}
						err := onr.Validate()
						require.Equal(v0ObjectValid, err == nil, "should be valid: %v %s", v0ObjectValid, err)

						asObject := &v0.RelationTuple{
							ObjectAndRelation: onr,
							User: &v0.User{
								UserOneof: &v0.User_Userset{
									Userset: knownGoodONR,
								},
							},
						}
						err = asObject.Validate()
						require.Equal(v0ObjectValid, err == nil, "should be valid: %v %s", v0ObjectValid, err)

						asSubject := &v0.RelationTuple{
							ObjectAndRelation: onr,
							User: &v0.User{
								UserOneof: &v0.User_Userset{
									Userset: &v0.ObjectAndRelation{
										Namespace: ns.name,
										ObjectId:  subjectID.name,
										Relation:  relation.name,
									},
								},
							},
						}
						err = asSubject.Validate()
						require.Equal(v0Valid, err == nil, "should be valid: %v %s", v0Valid, err)
					})
				}
			}
		}
	}
}

func TestV1CoreObjectValidity(t *testing.T) {
	for _, ns := range namespaces {
		for _, objectID := range objectIDs {
			for _, subjectID := range subjectIDs {
				for _, relation := range relations {
					testName := fmt.Sprintf("%s:%s#%s@%s", ns.name, objectID.name, relation.name, subjectID.name)
					t.Run(testName, func(t *testing.T) {
						t.Parallel()

						require := require.New(t)

						objRef := &v1.ObjectReference{
							ObjectType: ns.name,
							ObjectId:   objectID.name,
						}
						objRefValid := ns.valid && objectID.valid
						err := objRef.Validate()
						require.Equal(objRefValid, err == nil, "should be valid: %v %s", objRefValid, err)

						subjObjRef := &v1.ObjectReference{
							ObjectType: ns.name,
							ObjectId:   subjectID.name,
						}
						subObjRefValid := ns.valid && subjectID.valid
						err = subjObjRef.Validate()
						require.Equal(subObjRefValid, err == nil, "should be valid: %v %s", subObjRefValid, err)

						subRef := &v1.SubjectReference{
							Object:           subjObjRef,
							OptionalRelation: relation.name,
						}
						subjectValid := ns.valid && subjectID.valid &&
							(relation.validity == alwaysValid || relation.validity == validV1SubjectOnly)
						err = subRef.Validate()
						require.Equal(subjectValid, err == nil, "should be valid: %v %s", subjectValid, err)

						asResource := &v1.Relationship{
							Resource: objRef,
							Relation: relation.name,
							Subject:  knownGoodSubjectRef,
						}
						asResourceValid := objRefValid && relation.validity == alwaysValid
						err = asResource.Validate()
						require.Equal(asResourceValid, err == nil, "should be valid: %v %s", asResourceValid, err)

						asSubject := &v1.Relationship{
							Resource: knownGoodObjectRef,
							Relation: knownGoodSubjectRef.OptionalRelation,
							Subject:  subRef,
						}
						err = asSubject.Validate()
						require.Equal(subjectValid, err == nil, "should be valid: %v %s", subjectValid, err)

						// Test all the components of a filter
						justNS := &v1.RelationshipFilter{
							ResourceType: ns.name,
						}
						filterValid := ns.valid
						err = justNS.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						objectIDFilter := &v1.RelationshipFilter{
							ResourceType:       ns.name,
							OptionalResourceId: objectID.name,
						}
						filterValid = ns.valid && (objectID.valid || objectID.name == "")
						err = objectIDFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						objectRelationFilter := &v1.RelationshipFilter{
							ResourceType:     ns.name,
							OptionalRelation: relation.name,
						}
						filterValid = ns.valid && (relation.validity == alwaysValid || relation.name == "")
						err = objectRelationFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						fullObjectFilter := &v1.RelationshipFilter{
							ResourceType:       ns.name,
							OptionalResourceId: objectID.name,
							OptionalRelation:   relation.name,
						}
						filterValid = ns.valid && (objectID.valid || objectID.name == "") &&
							(relation.validity == alwaysValid || relation.name == "")
						err = fullObjectFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						bothTypesFilter := &v1.RelationshipFilter{
							ResourceType: knownGoodObjectRef.ObjectType,
							OptionalSubjectFilter: &v1.SubjectFilter{
								SubjectType: ns.name,
							},
						}
						filterValid = ns.valid
						err = bothTypesFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						subjectIDFilter := &v1.RelationshipFilter{
							ResourceType: knownGoodObjectRef.ObjectType,
							OptionalSubjectFilter: &v1.SubjectFilter{
								SubjectType:       ns.name,
								OptionalSubjectId: subjectID.name,
							},
						}
						filterValid = ns.valid && (subjectID.valid || subjectID.name == "")
						err = subjectIDFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						subjectRelationFilter := &v1.RelationshipFilter{
							ResourceType: knownGoodObjectRef.ObjectType,
							OptionalSubjectFilter: &v1.SubjectFilter{
								SubjectType: ns.name,
								OptionalRelation: &v1.SubjectFilter_RelationFilter{
									Relation: relation.name,
								},
							},
						}
						filterValid = ns.valid &&
							(relation.validity == alwaysValid || relation.validity == validV1SubjectOnly)
						err = subjectRelationFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)

						fullSubjectFilter := &v1.RelationshipFilter{
							ResourceType: knownGoodObjectRef.ObjectType,
							OptionalSubjectFilter: &v1.SubjectFilter{
								SubjectType:       ns.name,
								OptionalSubjectId: subjectID.name,
								OptionalRelation: &v1.SubjectFilter_RelationFilter{
									Relation: relation.name,
								},
							},
						}
						filterValid = ns.valid && (subjectID.valid || subjectID.name == "") &&
							(relation.validity == alwaysValid || relation.validity == validV1SubjectOnly)
						err = fullSubjectFilter.Validate()
						require.Equal(filterValid, err == nil, "should be valid: %v %s", filterValid, err)
					})
				}
			}
		}
	}
}

func TestV1CaveatValidity(t *testing.T) {
	for _, caveat := range caveats {
		testName := fmt.Sprintf("caveat->%s_context->%v", caveat.name, caveat.context)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			strct, err := structpb.NewStruct(caveat.context)
			require.NoError(err)

			optionalCaveat := &v1.ContextualizedCaveat{
				CaveatName: caveat.name,
				Context:    strct,
			}
			err = optionalCaveat.Validate()
			require.Equal(caveat.valid, err == nil, "should be valid: %v %s", caveat.valid, err)

			rel := &v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectType: "test",
					ObjectId:   "test",
				},
				Relation: "test",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: "test",
						ObjectId:   "test",
					},
				},
				OptionalCaveat: optionalCaveat,
			}
			err = rel.Validate()
			require.Equal(caveat.valid, err == nil, "should be valid: %v %s", caveat.valid, err)
		})
	}
}

func TestWildcardSubjectRelation(t *testing.T) {
	subjObjRef := &v1.ObjectReference{
		ObjectType: "somenamespace",
		ObjectId:   "*",
	}
	subRef := &v1.SubjectReference{
		Object:           subjObjRef,
		OptionalRelation: "somerelation",
	}
	require.Error(t, subRef.HandwrittenValidate())
}

func TestWildcardSubjectRelationEmpty(t *testing.T) {
	subjObjRef := &v1.ObjectReference{
		ObjectType: "somenamespace",
		ObjectId:   "*",
	}
	subRef := &v1.SubjectReference{
		Object: subjObjRef,
	}
	require.NoError(t, subRef.HandwrittenValidate())
}

func generateMap(length int) map[string]any {
	output := make(map[string]any, length)
	for i := 0; i < length; i++ {
		random := randString(32)
		output[random] = random
	}
	return output
}

var randInput = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = randInput[rand.Intn(len(randInput))] //nolint:gosec
	}
	return string(b)
}
