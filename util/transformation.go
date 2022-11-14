package util

import (
	"github.com/linguohua/titan/api"
	"sort"
)

// UniformMapping cid to many api.DownloadInfo mapping to cid to only one api.DownloadInfo
// the cid is evenly distributed on the edge nodes
func UniformMapping(cidToEdges map[string][]api.DownloadInfo) map[string]api.DownloadInfo {
	result := make(map[string]api.DownloadInfo)
	// api.DownloadInfo mapping to many cid
	edgeToCids := make(map[api.DownloadInfo]map[string]struct{})

	for key, downloadInfos := range cidToEdges {
		for _, downloadInfo := range downloadInfos {
			if cids, ok := edgeToCids[downloadInfo]; ok {
				cids[key] = struct{}{}
			} else {
				mid := make(map[string]struct{})
				mid[key] = struct{}{}
				edgeToCids[downloadInfo] = mid
			}
		}
	}

	// sort according to the number of cid contained in each edge node, asc sort
	sortDownloadInfo := make([]*downloadInfoCounter, 0, len(edgeToCids))
	for edge, cids := range edgeToCids {
		dc := &downloadInfoCounter{
			downloadInfo: edge,
			counter:      len(cids),
		}
		sortDownloadInfo = append(sortDownloadInfo, dc)
	}
	sort.Slice(sortDownloadInfo, func(i, j int) bool {
		return sortDownloadInfo[i].counter < sortDownloadInfo[j].counter
	})

	ring := newDownloadInfoRing(sortDownloadInfo)
	for ring.next() {
		if len(cidToEdges) == 0 {
			break
		}
		cv := ring.currentValue()
		cidStr := ""
		edgeCounter := 0
		for key := range edgeToCids[cv.downloadInfo] {
			if edges, ok := cidToEdges[key]; ok {
				if cidStr == "" {
					// first
					cidStr = key
					edgeCounter = len(edges)
				}
				if edgeCounter > len(edges) {
					cidStr = key
					edgeCounter = len(edges)
				}
			}
		}
		if cidStr == "" {
			continue
		}
		result[cidStr] = cv.downloadInfo
		// cid mapping edge, then delete the cid
		delete(cidToEdges, cidStr)
	}

	return result
}

type downloadInfoCounter struct {
	downloadInfo api.DownloadInfo
	counter      int
}

type downloadInfoRing struct {
	data  []*downloadInfoCounter
	index int
}

func newDownloadInfoRing(ds []*downloadInfoCounter) downloadInfoRing {
	return downloadInfoRing{
		data:  ds,
		index: -1,
	}
}

func (d *downloadInfoRing) next() bool {
	if d.data == nil || len(d.data) == 0 {
		return false
	}
	d.index++
	if d.index >= len(d.data) {
		d.index = 0
	}
	return true
}

func (d *downloadInfoRing) currentValue() *downloadInfoCounter {
	return d.data[d.index]
}
