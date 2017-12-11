package etl

import (
	"fmt"
	"os"
	"strings"

	"github.com/compose/transporter/adaptor"
	_ "github.com/compose/transporter/adaptor/all"
	_ "github.com/compose/transporter/function/all"
	"github.com/compose/transporter/log"

	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/data"
	"github.com/dailyburn/ratchet/logger"
	"github.com/dailyburn/ratchet/processors"
	"github.com/dailyburn/ratchet/util"
)

/*
	Refs:
	- github.com/compose/transporter
*/

type ETL interface {
}

type config struct {
	transporter *transporter.s
}
