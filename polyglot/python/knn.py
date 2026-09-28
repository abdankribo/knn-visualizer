import math

def classify(points, target, k=3):
    ranked=sorted(((math.dist((p[0],p[1]), target), p[2]) for p in points), key=lambda x:x[0])
    neighbors=ranked[:max(1,k)]
    votes={label:sum(1 for _,x in neighbors if x==label) for _,label in neighbors}
    return max(votes,key=votes.get), neighbors
