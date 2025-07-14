package node

type Lagu struct {
	ID       int
	Judul    string
	Penyanyi string
	Tahun    int
	Durasi   string
}

type NodeLaguBerlisensi struct {
	Data Lagu
	Next *NodeLaguBerlisensi
}

type NodeLaguTidakBerlisensi struct {
	IDEntry     int
	IDReferensi int
	Next        *NodeLaguTidakBerlisensi
}