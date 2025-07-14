package model

import "uts_strukturdata/node"

var LaguBerlisensiHead *node.NodeLaguBerlisensi
var LaguTidakBerlisensiHead *node.NodeLaguTidakBerlisensi

func TambahLaguBerlisensi(lg node.Lagu) {
	nodeLagu := &node.NodeLaguBerlisensi{Data: lg}
	if LaguBerlisensiHead == nil {
		LaguBerlisensiHead = nodeLagu
	} else {
		temp := LaguBerlisensiHead
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = nodeLagu
	}
}

func CariLaguBerlisensi(id int) *node.Lagu {
	temp := LaguBerlisensiHead
	for temp != nil {
		if temp.Data.ID == id {
			return &temp.Data
		}
		temp = temp.Next
	}
	return nil
}

func TambahLaguTidakBerlisensi(idEntry int, idReferensi int) bool {
	if CariLaguBerlisensi(idReferensi) == nil {
		return false
	}
	nodeLagu := &node.NodeLaguTidakBerlisensi{
		IDEntry:     idEntry,
		IDReferensi: idReferensi,
	}
	if LaguTidakBerlisensiHead == nil {
		LaguTidakBerlisensiHead = nodeLagu
	} else {
		temp := LaguTidakBerlisensiHead
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = nodeLagu
	}
	return true
}

func GetLaguTidakBerlisensi() *node.NodeLaguTidakBerlisensi {
	return LaguTidakBerlisensiHead
}
