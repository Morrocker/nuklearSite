package main

import (
	"github.com/morrocker/log"
	"github.com/morrocker/recoveryserver/recovery2/files"
	"github.com/morrocker/recoveryserver/recovery2/remote"
	track "github.com/morrocker/recoveryserver/recovery2/tracker"
	"github.com/morrocker/recoveryserver/recovery2/tree"
)

func main() {
	data := tree.Data{
		RootId:     "3dbdd742a408d242b03e18d1d6299f1e6e61dd733de68b4018d89e7d53fe8652b15c78b4026303e49c25f90d1b5f8355a6077684777eaf520470fff2a42ad539",
		Repository: "382e0bba1bc3f122b88f41a6b87483d2feb43ce3f15ff71bd600b1b2077ac64d3cf2f6036f0297cc77667baf270cce7bf077f640391a30f56abcbb279c5e429b",
		Server:     "https://zifre.cloner.cl/",
		ClonerKey:  "ed9a015afd459c97f2e24cb0d809e440e8cafe8d02979d04e605cef66fe806b09bcf3cd0dd788c88a979e88a885b229b9bcf99c2f61779acc0bdb150c3ffcacc",
		Version:    999999999999,
		Deleted:    false,
		Exclusions: make(map[string]bool),
	}

	thrott := tree.Throttling{
		BuffSize: 10000,
		Workers:  2,
	}

	tr := track.New()

	mt, err := tree.GetRecoveryTree(data, thrott, tr)
	if err != nil {
		log.Errorln(err)
	}

	fdata := files.Data{
		User:    "maximo.ormeno@arysta.com",
		Legacy:  false,
		Workers: 30,
	}

	rbs := remote.NewRBS("https://alpha.cloner.cl:4000", "youarethebest")

	if err := files.GetFiles(mt, "./", fdata, rbs, tr); err != nil {
		log.Errorln(err)
	}
}
