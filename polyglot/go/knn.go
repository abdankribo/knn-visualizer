package knn
import "math"

type Point struct{ X,Y float64; Label string }
func Classify(p []Point, tx,ty float64,k int) string { counts:=map[string]int{}; for i:=0;i<len(p);i++ { best:=i; for j:=i+1;j<len(p);j++ { if math.Hypot(p[j].X-tx,p[j].Y-ty)<math.Hypot(p[best].X-tx,p[best].Y-ty){best=j} }; p[i],p[best]=p[best],p[i] }; for i:=0;i<k&&i<len(p);i++{counts[p[i].Label]++}; if counts["A"]>=counts["B"] {return "A"}; return "B" }
