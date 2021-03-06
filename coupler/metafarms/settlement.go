package metafarms

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/agamigo/agamigo/coupler"
)

type Killsheet struct {
	plant            string
	producer         string
	producerID       int
	farm             string
	location         string
	tattoo           string
	purchaseDate     time.Time
	killDate         time.Time
	shift            int
	head             int
	deadCount        int
	condemndCount    int
	liveWeight       float64
	liveWeightAvg    float64
	liveWeightAdj    float64
	carcassWeight    float64
	carcassWeightAvg float64
	carcassYieldAvg  float64
	percentLeanAvg   float64
	sortAdj          float64
	sortAdjAvg       float64
	leanAdj          float64
	leanAdjAvg       float64
	leanPctAvg       float64
	backfatAvg       float64
	loinDepthAvg     float64
	uniformLeanStats map[string]float64
	carcassBasePrice float64
	carcassValueCWT  float64
	carcassValue     float64
	liveValueAdjCWT  float64
	grossAmount      float64
	expenses         map[string]float64
	netAmount        float64
	checkNumber      int
	checkDate        time.Time
	weightGroups     []*weightGroup
}

type weightGroup struct {
	weightRange *coupler.FloatRange
	leanRange   *coupler.FloatRange
	head        int
	weight      float64
	basePrice   float64
	sortAdj     float64
	leanAdj     float64
	value       float64
}

func NewKillsheetsFromCSV(r io.Reader) (kss []*Killsheet, err error) {
	cr, err := newCSVReader(r)
	if err != nil {
		return nil, fmt.Errorf("Unable to sanitize CSV data: %v", err)
	}

	cr.FieldsPerRecord = -1
	ks := &Killsheet{
		uniformLeanStats: map[string]float64{},
		expenses:         map[string]float64{},
	}

	for {
		line, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return nil, err
		}

		done, err := ks.parseLine(line)
		if err != nil {
			// log.Fatalf("Issue with line: %#v\n", line)
			return nil, err
		}

		if done {
			kss = append(kss, ks)
			ks = &Killsheet{}
		}
	}

	return kss, nil
}

func newCSVReader(r io.Reader) (*csv.Reader, error) {
	p, err := ioutil.ReadAll(r)
	if err != nil {
		return &csv.Reader{}, fmt.Errorf("Unable to remove CR from reader: %v", err)
	}

	ps := bytes.Split(p, []byte("\n"))
	for i, _ := range ps {
		ps[i] = bytes.Trim(ps[i], " \r")
	}
	p = bytes.Join(ps, []byte("\n"))

	return csv.NewReader(bytes.NewReader(p)), nil
}

func (ks *Killsheet) parseLine(l []string) (done bool, err error) {
	for i, _ := range l {
		l[i] = strings.TrimSpace(l[i])
	}

	switch s := l[0]; s {
	case "HEADER01":
		ks.plant = l[1]
		ks.producerID, err = strconv.Atoi(l[2])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Producer ID: %v", err)
		}
		ks.farm = l[3]
		ks.location = l[4]
		ks.producer = l[5]

		return done, nil
	case "HEADER02":
		ks.tattoo = l[1]
		ks.purchaseDate, err = time.Parse("01/02/06", l[2])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Purchase Date: %v", err)
		}
		ks.killDate, err = time.Parse("01/02/06", l[3])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Kill Date: %v", err)
		}
		ks.shift, err = strconv.Atoi(l[4])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Shift: %v", err)
		}
		ks.head, err = strconv.Atoi(l[5])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Head: %v", err)
		}
		ks.carcassBasePrice, err = strconv.ParseFloat(l[6], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Carcass Base Price: %v", err)
		}
		ks.liveWeight, err = strconv.ParseFloat(l[7], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Live Weight: %v", err)
		}
		ks.deadCount, err = strconv.Atoi(l[8])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Dead Count: %v", err)
		}
		ks.condemndCount, err = strconv.Atoi(l[9])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Condemnd Count: %v", err)
		}
		return done, err
	case "DETAIL":
		wg := &weightGroup{}

		// TODO: Figure out how ADJUSTMENT fits into this
		wg.weightRange, err = coupler.Ator(l[1])
		if err != nil && l[1] != "ADJUSTMENT" {
			return false, fmt.Errorf("Unable to parse Weight Range: %v", err)
		}

		wg.leanRange, err = coupler.Ator(l[2])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Pct Lean Range: %v", err)
		}

		wg.head, err = strconv.Atoi(l[3])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Group Head: %v", err)
		}

		wg.weight, err = strconv.ParseFloat(l[4], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Group Weight: %v", err)
		}

		wg.basePrice, err = strconv.ParseFloat(l[5], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Group Base Price: %v", err)
		}

		wg.sortAdj, err = strconv.ParseFloat(l[6], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Group Sort Adj: %v", err)
		}

		wg.leanAdj, err = strconv.ParseFloat(l[7], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Group Lean Premium: %v", err)
		}

		wg.value, err = strconv.ParseFloat(l[8], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Group Carcass Value: %v", err)
		}

		ks.weightGroups = append(ks.weightGroups, wg)
		return false, err
	case "TOTAL01":
		// TODO: l[1] redundant head count, maybe use for sanity check
		ks.carcassWeight, err = strconv.ParseFloat(l[2], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Carcass Weight: %v", err)
		}

		// TODO: l[3] unknown

		ks.sortAdj, err = strconv.ParseFloat(l[4], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Total Sort Adj: %v", err)
		}

		ks.leanAdj, err = strconv.ParseFloat(l[5], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Total Lean Adj: %v", err)
		}

		ks.carcassValue, err = strconv.ParseFloat(l[6], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Total Carcass Value: %v", err)
		}

		return false, nil
	case "TOTAL02":
		// TODO: Use as sanity check against weight groups
		// TOTAL0{2,3,4,5} are Pct Lean groups. Redundant.
		return false, nil
	case "TOTAL03":
		// TOTAL0{2,3,4,5} are Pct Lean groups. Redundant.
		return false, nil
	case "TOTAL04":
		// TOTAL0{2,3,4,5} are Pct Lean groups. Redundant.
		return false, nil
	case "TOTAL05":
		// TOTAL0{2,3,4,5} are Pct Lean groups. Redundant.
		return false, nil
	case "TOTAL06":
		ks.uniformLeanStats["ffli"], err = strconv.ParseFloat(l[1], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Uniform Lean FFLI: %v", err)
		}

		ks.uniformLeanStats["bf"], err = strconv.ParseFloat(l[2], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Uniform Lean BF: %v", err)
		}

		ks.uniformLeanStats["hcw"], err = strconv.ParseFloat(l[3], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Uniform Lean HCW: %v", err)
		}

		return false, nil
	case "FINAL01":
		// TODO: l[1] redundant carcass weight, maybe use for sanity check
		// TODO: l[2] redundant carcass value, maybe use for sanity check
		// TODO: l[3] redundant base price, maybe use for sanity check
		ks.sortAdjAvg, err = strconv.ParseFloat(l[4], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Sort Adjust Avg: %v", err)
		}

		ks.leanAdjAvg, err = strconv.ParseFloat(l[5], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Lean Adjust Avg: %v", err)
		}

		ks.carcassValueCWT, err = strconv.ParseFloat(l[6], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Carcass Value CWT: %v", err)
		}

		ks.leanPctAvg, err = strconv.ParseFloat(l[7], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Lean Percent Avg: %v", err)
		}
		return false, nil
	case "FINAL02":
		ks.backfatAvg, err = strconv.ParseFloat(l[1], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Back Fat Avg: %v", err)
		}

		ks.loinDepthAvg, err = strconv.ParseFloat(l[2], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Loin Depth Avg: %v", err)
		}

		ks.liveWeightAvg, err = strconv.ParseFloat(l[3], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Live Weight Avg: %v", err)
		}

		ks.carcassWeightAvg, err = strconv.ParseFloat(l[4], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Carcass Weight Avg: %v", err)
		}

		ks.liveWeightAdj, err = strconv.ParseFloat(l[5], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Live Weight Adjusted: %v", err)
		}

		ks.carcassYieldAvg, err = strconv.ParseFloat(l[6], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Shift: %v", err)
		}

		ks.liveValueAdjCWT, err = strconv.ParseFloat(l[7], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Shift: %v", err)
		}

		return false, nil
	case "FINAL03":
		ks.grossAmount, err = strconv.ParseFloat(l[1], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Gross Value: %v", err)
		}

		ks.expenses["pork_board"], err = strconv.ParseFloat(l[2], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Pork Board Expense: %v", err)
		}

		// TODO: l[3,4,5,6] contain unknown expenses/fees

		ks.netAmount, err = strconv.ParseFloat(l[7], 64)
		if err != nil {
			return false, fmt.Errorf("Unable to parse Net Value: %v", err)
		}

		ks.checkNumber, err = strconv.Atoi(l[8])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Check Number: %v", err)
		}

		ks.checkDate, err = time.Parse("01/02/06", l[9])
		if err != nil {
			return false, fmt.Errorf("Unable to parse Check Date: %v", err)
		}

		return true, nil
	}

	return false, errors.Errorf("Unable to parse line into killsheet.")
}
