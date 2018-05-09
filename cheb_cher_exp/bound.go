package cheb_cher_exp

import "math"

type bounds struct {
	aa []float64
	a uint64
}
func Newbounds(aa []float64,a uint64) bounds{
	b:=bounds{
		aa:aa,
		a:a,
	}
	return b

}
func (b bounds)Experimantal()  (float64,float64){
	succes:=0
	upper:=0.0
	lower:=0.0
	for true{
		succes=0
		for i:=0;i<len(b.aa);i++{
			if b.aa[i]<=1+upper && b.aa[i]>=1+lower {
				succes++
		}
		}
		if float64(succes)/float64(len(b.aa))>=float64(1-(b.a/100)){
			break
		} else{
			upper+=0.001
			lower-=0.001
		}
	}



	return upper, lower


}

func (b bounds)Chebyshev(n,k uint64) float64 {
	variancja:=float64(1)/float64(2*k)
	bv:=0.0
	result:=0.0
	for true{
		result=1-(variancja/math.Pow(bv,2))
		if result>=1.0-(float64(b.a)/100){
			break
		}else{
			bv+=0.001
		}
	}

	return bv
}
