package titles

import (
	"sort"
)

func (cmp ResultComparer) Sort(res Results) {
	sorter := &resultSorter{
		res: res,
		by:  cmp,
	}
	sort.Sort(sorter)
}

type resultSorter struct {
	by  ResultComparer
	res Results
}

func (f *resultSorter) Len() int {
	return len(f.res)
}

func (f *resultSorter) Less(i, j int) bool {
	return f.by(&f.res[i], &f.res[j])
}

func (f *resultSorter) Swap(i, j int) {
	f.res[i], f.res[j] = f.res[j], f.res[i]
}
