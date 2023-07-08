package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var factors = []struct {
	input    int
	expected []int
}{
	{2, []int{2}},
	{3, []int{3}},
	{4, []int{2, 2}},
	{5, []int{5}},
	{6, []int{2, 3}},
	{7, []int{7}},
	{8, []int{2, 2, 2}},
	{9, []int{3, 3}},
	{10, []int{2, 5}},
	{100, []int{2, 2, 5, 5}},
	{1000, []int{2, 2, 2, 5, 5, 5}},
}

func TestPrimeFactorsSpecialCase1(t *testing.T) {
	assert.Equal(t, []int{}, primeFactors(1))
}

func TestPrimeFactors(t *testing.T) {
	for _, tt := range factors {
		assert.Equal(t, tt.expected, primeFactors(tt.input))
	}
}
