// Code generated by argen. DO NOT EDIT.
// This code was generated from a template.
//
// Manual changes to this file may cause unexpected behavior in your application.
// Manual changes to this file will be overwritten if the code is regenerated.
//
// Generate info: argen@v1.5.3-18-g3247b15 (Commit: 3247b15e)
package repository

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/arobj"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/boolindexed"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/promoperiods"
	"github.com/mailru/activerecord-cookbook/example/model/repository/generated/reward"
	"github.com/mailru/activerecord/pkg/octopus"
)

type SpaceMeta struct {
	PackageName string
	Unpacker    func(ctx context.Context, tuple octopus.TupleData) (any, error)
	Fields      []FieldMeta
	PK          IndexMeta
	Indexes     map[string]IndexMeta
}

type IndexMeta struct {
	Name     string
	Unpacker func(packedKeys [][][]byte) (any, error)
}

type FieldMeta struct {
	Name     string
	Unpacker func(packedField []byte) (any, error)
}

type NSPackage map[string]SpaceMeta

func (ns NSPackage) meta(n uint32) (SpaceMeta, bool) {
	v, ok := ns[strconv.Itoa(int(n))]
	return v, ok
}

var NamespacePackages = NSPackage{

	"5": {
		PackageName: "arobj",
		Unpacker: func(ctx context.Context, tuple octopus.TupleData) (any, error) {
			obj, err := arobj.TupleToStruct(ctx, tuple)
			if err != nil {
				return nil, fmt.Errorf("can't decode tuple: %s", err)
			}

			return arobj.MarshalFixtures([]*arobj.ArObj{obj})
		},
		Fields: []FieldMeta{
			{
				Name:     "ID",
				Unpacker: func(packedField []byte) (any, error) { return arobj.UnpackID(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Name",
				Unpacker: func(packedField []byte) (any, error) { return arobj.UnpackName(bytes.NewReader(packedField)) },
			},
			{
				Name:     "AnotherID",
				Unpacker: func(packedField []byte) (any, error) { return arobj.UnpackAnotherID(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Type",
				Unpacker: func(packedField []byte) (any, error) { return arobj.UnpackType(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Flags",
				Unpacker: func(packedField []byte) (any, error) { return arobj.UnpackFlags(bytes.NewReader(packedField)) },
			},
		},
		Indexes: map[string]IndexMeta{

			"0.1": {
				Name:     "ID",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return arobj.UnpackKeyIndexID(packedKeys) },
			},
			"1.1": {
				Name:     "Type",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return arobj.UnpackKeyIndexType(packedKeys) },
			},
			"2.2": {
				Name:     "TypeID",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return arobj.UnpackKeyIndexTypeID(packedKeys) },
			},
			"2.1": {
				Name:     "TypePart",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return arobj.UnpackKeyIndexTypePart(packedKeys) },
			},
		},
		PK: IndexMeta{

			Name:     "ID",
			Unpacker: func(packedKeys [][][]byte) (any, error) { return arobj.UnpackKeyIndexID(packedKeys) },
		},
	},

	"25": {
		PackageName: "boolindexed",
		Unpacker: func(ctx context.Context, tuple octopus.TupleData) (any, error) {
			obj, err := boolindexed.TupleToStruct(ctx, tuple)
			if err != nil {
				return nil, fmt.Errorf("can't decode tuple: %s", err)
			}

			return boolindexed.MarshalFixtures([]*boolindexed.Boolindexed{obj})
		},
		Fields: []FieldMeta{
			{
				Name:     "Code",
				Unpacker: func(packedField []byte) (any, error) { return boolindexed.UnpackCode(bytes.NewReader(packedField)) },
			},
			{
				Name: "Invisible",
				Unpacker: func(packedField []byte) (any, error) {
					return boolindexed.UnpackInvisible(bytes.NewReader(packedField))
				},
			},
		},
		Indexes: map[string]IndexMeta{

			"0.1": {
				Name:     "Code",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return boolindexed.UnpackKeyIndexCode(packedKeys) },
			},
			"1.1": {
				Name:     "Invisible",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return boolindexed.UnpackKeyIndexInvisible(packedKeys) },
			},
		},
		PK: IndexMeta{

			Name:     "Code",
			Unpacker: func(packedKeys [][][]byte) (any, error) { return boolindexed.UnpackKeyIndexCode(packedKeys) },
		},
	},

	"6": {
		PackageName: "promoperiods",
		Unpacker: func(ctx context.Context, tuple octopus.TupleData) (any, error) {
			obj, err := promoperiods.TupleToStruct(ctx, tuple)
			if err != nil {
				return nil, fmt.Errorf("can't decode tuple: %s", err)
			}

			return promoperiods.MarshalFixtures([]*promoperiods.Promoperiods{obj})
		},
		Fields: []FieldMeta{
			{
				Name:     "ID",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackID(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Code",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackCode(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Email",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackEmail(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Start",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackStart(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Finish",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackFinish(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Action",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackAction(bytes.NewReader(packedField)) },
			},
			{
				Name: "Platform",
				Unpacker: func(packedField []byte) (any, error) {
					return promoperiods.UnpackPlatform(bytes.NewReader(packedField))
				},
			},
			{
				Name: "Promobunch",
				Unpacker: func(packedField []byte) (any, error) {
					return promoperiods.UnpackPromobunch(bytes.NewReader(packedField))
				},
			},
			{
				Name: "Platforms",
				Unpacker: func(packedField []byte) (any, error) {
					return promoperiods.UnpackPlatforms(bytes.NewReader(packedField))
				},
			},
			{
				Name:     "PlanID",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackPlanID(bytes.NewReader(packedField)) },
			},
			{
				Name: "PlanType",
				Unpacker: func(packedField []byte) (any, error) {
					return promoperiods.UnpackPlanType(bytes.NewReader(packedField))
				},
			},
			{
				Name:     "Price",
				Unpacker: func(packedField []byte) (any, error) { return promoperiods.UnpackPrice(bytes.NewReader(packedField)) },
			},
		},
		Indexes: map[string]IndexMeta{

			"0.1": {
				Name:     "ID",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexID(packedKeys) },
			},
			"1.1": {
				Name:     "Code",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexCode(packedKeys) },
			},
			"2.1": {
				Name:     "Email",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexEmail(packedKeys) },
			},
			"3.2": {
				Name:     "PlanTypePrice",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexPlanTypePrice(packedKeys) },
			},
			"4.2": {
				Name:     "EmailCode",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexEmailCode(packedKeys) },
			},
			"5.2": {
				Name:     "EmailAction",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexEmailAction(packedKeys) },
			},
			"4.1": {
				Name:     "EmailPart",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexEmailPart(packedKeys) },
			},
		},
		PK: IndexMeta{

			Name:     "ID",
			Unpacker: func(packedKeys [][][]byte) (any, error) { return promoperiods.UnpackKeyIndexID(packedKeys) },
		},
	},

	"24": {
		PackageName: "reward",
		Unpacker: func(ctx context.Context, tuple octopus.TupleData) (any, error) {
			obj, err := reward.TupleToStruct(ctx, tuple)
			if err != nil {
				return nil, fmt.Errorf("can't decode tuple: %s", err)
			}

			return reward.MarshalFixtures([]*reward.Reward{obj})
		},
		Fields: []FieldMeta{
			{
				Name:     "Code",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackCode(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Services",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackServices(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Partner",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackPartner(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Extra",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackExtra(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Flags",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackFlags(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Unlocked",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackUnlocked(bytes.NewReader(packedField)) },
			},
			{
				Name:     "Description",
				Unpacker: func(packedField []byte) (any, error) { return reward.UnpackDescription(bytes.NewReader(packedField)) },
			},
		},
		Indexes: map[string]IndexMeta{

			"0.1": {
				Name:     "Code",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return reward.UnpackKeyIndexCode(packedKeys) },
			},
			"1.1": {
				Name:     "Partner",
				Unpacker: func(packedKeys [][][]byte) (any, error) { return reward.UnpackKeyIndexPartner(packedKeys) },
			},
		},
		PK: IndexMeta{

			Name:     "Code",
			Unpacker: func(packedKeys [][][]byte) (any, error) { return reward.UnpackKeyIndexCode(packedKeys) },
		},
	},
}

func (n NSPackage) GetSelectDebugInfo(ns uint32, indexnum uint32, offset uint32, limit uint32, keys [][][]byte) string {
	spacemeta, ex := n.meta(ns)
	if !ex {
		return fmt.Sprintf("unknown space %d, index: %d, offset: %d, limit: %d, Keys: %+v", ns, indexnum, offset, limit, keys)
	}

	ind, ex := spacemeta.Indexes[fmt.Sprintf("%d.%d", indexnum, len(keys[0]))]
	if !ex {
		return fmt.Sprintf("space %d (%s), unknown index: %d (%d.%d), offset: %d, limit: %d, Keys: %+v", ns, spacemeta.PackageName, indexnum, indexnum, len(keys[0]), offset, limit, keys)
	}

	unpackedKeys, err := ind.Unpacker(keys)
	if err != nil {
		return fmt.Sprintf("Space: %d (%s), index: %d (%s), offset: %d, limit: %d, Keys: %+v (error unpack: %s)", ns, spacemeta.PackageName, indexnum, ind.Name, offset, limit, keys, err)
	}

	return fmt.Sprintf("Space: %d (%s), index: %d (%s), offset: %d, limit: %d, Keys: %+v", ns, spacemeta.PackageName, indexnum, ind.Name, offset, limit, unpackedKeys)
}

func (n NSPackage) GetUpdateDebugInfo(ns uint32, primaryKey [][]byte, updateOps []octopus.Ops) string {
	spacemeta, ex := n.meta(ns)
	if !ex {
		return fmt.Sprintf("unknown space %d, primaryKey: %+v, updateOps: %+v", ns, primaryKey, updateOps)
	}

	unpackedKeys, err := spacemeta.PK.Unpacker([][][]byte{primaryKey})
	if err != nil {
		return fmt.Sprintf("Space: %d (%s), primaryKey: %+v, updateOps: %+v (error unpack: %s)", ns, spacemeta.PackageName, primaryKey, updateOps, err)
	}

	updateFields := ""

	for _, op := range updateOps {
		val, err := spacemeta.Fields[op.Field].Unpacker(op.Value)
		if err != nil {
			val = fmt.Sprintf("% X (can't unpack: %s)", op.Value, err)
		}

		updateFields += fmt.Sprintf("%s %s <= `%v`; ", octopus.GetOpCodeName(op.Op), spacemeta.Fields[op.Field].Name, val)
	}

	return fmt.Sprintf("Space: %d (%s), primaryKey: %s (%+v), updateOps: %s", ns, spacemeta.PackageName, spacemeta.PK.Name, unpackedKeys, updateFields)
}

func (n NSPackage) GetDeleteDebugInfo(ns uint32, primaryKey [][]byte) string {
	spacemeta, ex := n.meta(ns)
	if !ex {
		return fmt.Sprintf("unknown space %d, primaryKey: %+v", ns, primaryKey)
	}

	unpackedKeys, err := spacemeta.PK.Unpacker([][][]byte{primaryKey})
	if err != nil {
		return fmt.Sprintf("Space: %d (%s), primaryKey: %+v (error unpack: %s)", ns, spacemeta.PackageName, primaryKey, err)
	}

	return fmt.Sprintf("Space: %d (%s), primaryKey: %s (%+v)", ns, spacemeta.PackageName, spacemeta.PK.Name, unpackedKeys)
}

func (n NSPackage) GetInsertDebugInfo(ns uint32, needRetVal bool, insertMode octopus.InsertMode, tuple octopus.TupleData) string {
	strMode := octopus.GetInsertModeName(insertMode)

	spacemeta, ex := n.meta(ns)
	if !ex {
		return fmt.Sprintf("unknown space %d, insertMode: %s, tuple: %+v", ns, strMode, tuple)
	}

	strObj, err := spacemeta.Unpacker(context.TODO(), tuple)
	if err != nil {
		return fmt.Sprintf("Space: %d (%s), insertMode: %s, tuple: %+v (err unpack: %s)", ns, spacemeta.PackageName, strMode, tuple, err)
	}

	return fmt.Sprintf("Space: %d (%s), insertMode: %s, tuple: \n%s", ns, spacemeta.PackageName, strMode, strObj)
}
