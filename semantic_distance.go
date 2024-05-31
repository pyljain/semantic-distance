package semanticdistance

import "math"

type Algorithm string

const (
	AlgorithmCosine Algorithm = "cosine"
	AlgorithmL2     Algorithm = "L2"
)

func Distance(distanceType Algorithm, searchVector []float64, corpus [][]float64) []float64 {
	var results []float64
	for i := 0; i < len(corpus); i++ {
		if distanceType == AlgorithmCosine {
			results = append(results, cosineSimilarity(searchVector, corpus[i]))
		} else if distanceType == AlgorithmL2 {
			results = append(results, l2Distance(searchVector, corpus[i]))
		}
	}

	return results
}

func dotProduct(A, B []float64) float64 {
	dot := 0.0
	for i := 0; i < len(A); i++ {
		dot += A[i] * B[i]
	}
	return dot
}

func magnitude(A []float64) float64 {
	sum := 0.0
	for _, v := range A {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func cosineSimilarity(A, B []float64) float64 {
	dot := dotProduct(A, B)
	magA := magnitude(A)
	magB := magnitude(B)
	if magA == 0 || magB == 0 {
		return 0 // handle zero vector
	}
	return dot / (magA * magB)
}

func l2Distance(A, B []float64) float64 {
	var sum float64
	for i := 0; i < len(A); i++ {
		sum += (A[i] - B[i]) * (A[i] - B[i])
	}
	return math.Sqrt(sum)
}
