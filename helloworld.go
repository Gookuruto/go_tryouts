package main

import ("fmt"
	"encoding/binary"
	"bytes"
	"math"
	//"crypto/sha256"
	"sort"
	"time"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
	"hash/fnv"
	"awesomeProject/cheb_cher_exp"
)

func uint_tobytearr(x uint32) []byte{
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint32(bs, x)//change int to byte array
	//fmt.Println(bs)
	return bs
}
func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
func sha256_int_to_flt(y uint32)float64  {
	return  float64(y)
}

func kMin(k int,M[] int32) float64 {
	Arr := make([]float64, k)
	bs := make([]byte, 4)
	z:=fnv.New32()
	for index := range Arr {
		Arr[index] = 1.0
	}
	//var y int32
	for i := range M {
		bs = uint_tobytearr(uint32(M[i]))
		z.Write(bs)
		hashed :=z.Sum32()
		//fmt.Println(hashed)
		temp:= sha256_int_to_flt(hashed)/math.Pow(2,32)
		//fmt.Println(temp)
		if temp <= Arr[k-1] {
			Arr[k-1] = temp
			//fmt.Println(Arr[k-1])
			sort.Float64s(Arr)
		}
		//z.Reset()

	}
	if Arr[k-1] == 1{
		n := 0
		for ind := range Arr {
			if Arr[ind] < 1 {
				n++
			}
		}
		//fmt.Println(Arr[k-1])
		return float64(n)/float64(len(M))
	} else {
		x := (float64(k - 1)) / Arr[k-1]
		//fmt.Println(Arr[k-1])
		return x/float64(len(M))
	}
}

func bytearr_to_int(x []byte) float64{
	var y float64
	err := binary.Read(bytes.NewReader(x), binary.LittleEndian, &y)
	if err != nil {
		panic(err)
	}
	//fmt.Println(y)
	return math.Abs(y)
}
func generate_M(size,firstelement int32) []int32  {
	M :=make([]int32,0,0)
	j:=firstelement
	for i:=0;i<int(size);i++{
		M=append(M,j)
		j++


	}
return M
}

func genrate_many_M(how_many int32) [][]int32  {
	x:=make([][]int32,how_many,how_many)
	for i:=0;int32(i)<how_many;i++{
		if i==0 {
			//fmt.Println(i)
			x[i] = make([]int32, i+1)
			x[i]=generate_M(int32(i+1),1)
		}else{
			x[i] = make([]int32, i+1)
			x[i]=generate_M(int32(i+1),x[i-1][len(x[i-1])-1]+1)
		}

	}
	return x

}

func main() {
	start := time.Now()
	numbers:=genrate_many_M(10000)
	nn:=make([]float64,1000,1000)
	nn1:=make([]float64,1000,1000)
	nn2:=make([]float64,1000,1000)
	nn3:=make([]float64,1000,1000)
	nn4:=make([]float64,1000,1000)
	nn5:=make([]float64,1000,1000)
	n1:=make([]float64,1000,1000)
	//h:=sha256.Sum256(numbers)
	//h2:=h[:]
	for j:=0;j<1000;j++{
		n1[j]=float64(j+1)
	}
	for i:=0;i<1000;i++ {
		nn1[i]=kMin(2, numbers[i])
	}
	for i:=0;i<1000;i++ {
		nn2[i]=kMin(3, numbers[i])
	}
	for i:=0;i<1000;i++ {
		nn3[i]=kMin(100, numbers[i])
	}
	for i:=0;i<1000;i++  {
		nn4[i]=kMin(100, numbers[i])
	}
	for i:=0;i<1000;i++ {
		nn[i]=kMin(400, numbers[i])
	}
	succes:=0
	for k:=2;k<=20000;k+=2{
		succes=0
		for i:=0;i<1000;i++ {
			nn5[i]=kMin(k, numbers[i])
			if nn5[i]<=1.1 && nn5[i]>=0.9{
				succes++
			}
		}
		//fmt.Println(float64(succes)/1000)
		if float64(succes)/1000>=0.95{
			fmt.Println("optymalne k wynosi",k)//952
			break
		}
	}
	x:=cheb_cher_exp.Newbounds(nn,5)
	fmt.Println(x.Experimantal())
	fmt.Println(x.Chebyshev(10000,400))
	pltor:=make(plotter.XYs,len(nn3))
	for i:=0;i<len(nn3);i++{
		pltor[i].X=float64(i+1)
		pltor[i].Y=nn3[i]


	}
	p,err:=plot.New()
	if err!=nil{
		panic(err)
	}
	p.Title.Text = "HyperLogLog"
	p.X.Label.Text="n"
	p.Y.Label.Text="estimator/n"
	err=plotutil.AddScatters(p,"First",pltor)
	//Save plot to png file
	if err :=p.Save(10*vg.Inch,10*vg.Inch,"estimation1.png");err!=nil{
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
