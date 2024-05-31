package semanticdistance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCosineSimilarity(t *testing.T) {
	A := []float64{1, 2, 3}
	B := []float64{4, 5, 6}
	C := []float64{7, 8, 9}
	D := []float64{10, 11, 12}

	if cosineSimilarity(A, B) != 0.9746318461970762 {
		t.Errorf("cosineSimilarity(A, B) = %v, want %v", cosineSimilarity(A, B), 0.9746318461970762)
	}
	if cosineSimilarity(A, C) != 0.9594119455666703 {
		t.Errorf("cosineSimilarity(A, C) = %v, want %v", cosineSimilarity(A, C), 0.9594119455666703)
	}
	if cosineSimilarity(A, D) != 0.9512583076673059 {
		t.Errorf("cosineSimilarity(A, D) = %v, want %v", cosineSimilarity(A, D), 0.9512583076673059)
	}
}

func TestL2Distance(t *testing.T) {
	A := []float64{1, 2, 3}
	B := []float64{4, 5, 6}
	C := []float64{7, 8, 9}
	D := []float64{10, 11, 12}

	if l2Distance(A, B) != 5.196152422706632 {
		t.Errorf("l2Distance(A, B) = %v, want %v", l2Distance(A, B), 5.196152422706632)
	}
	if l2Distance(A, C) != 10.392304845413264 {
		t.Errorf("l2Distance(A, C) = %v, want %v", l2Distance(A, C), 10.392304845413264)
	}
	if l2Distance(A, D) != 15.588457268119896 {
		t.Errorf("l2Distance(A, D) = %v, want %v", l2Distance(A, D), 15.588457268119896)
	}
}

func TestDistance(t *testing.T) {
	A := []float64{1, 2, 3}
	B := []float64{4, 5, 6}
	C := []float64{7, 8, 9}
	D := []float64{10, 11, 12}

	require.Equal(t, []float64{0.9746318461970762, 0.9594119455666703, 0.9512583076673059}, Distance(AlgorithmCosine, A, [][]float64{B, C, D}))
	require.Equal(t, []float64{5.196152422706632, 10.392304845413264, 15.588457268119896}, Distance(AlgorithmL2, A, [][]float64{B, C, D}))
}
