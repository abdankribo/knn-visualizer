#include <algorithm>
#include <cmath>
#include <string>
#include <vector>
struct Point{double x,y;std::string label;};
std::string classify(std::vector<Point> p,double tx,double ty,int k){std::sort(p.begin(),p.end(),[&](auto&a,auto&b){return std::hypot(a.x-tx,a.y-ty)<std::hypot(b.x-tx,b.y-ty);});int a=0,b=0;for(int i=0;i<k&&i<(int)p.size();++i)(p[i].label=="A"?a:b)++;return a>=b?"A":"B";}
