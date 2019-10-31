package example

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/itsubaki/gostream/pkg/builder"
	"github.com/itsubaki/gostream/pkg/event"
	"github.com/itsubaki/gostream/pkg/function"
	"github.com/itsubaki/gostream/pkg/parser"
	"github.com/itsubaki/gostream/pkg/selector"
	"github.com/itsubaki/gostream/pkg/view"
	"github.com/itsubaki/gostream/pkg/window"
)

func TimeWindow() {
	type LogEvent struct {
		Time    time.Time
		Level   int
		Message string
	}

	w := window.NewTime(LogEvent{}, 10*time.Second)
	defer w.Close()

	w.SetSelector(
		selector.LargerThanInt{
			Name:  "Level",
			Value: 2,
		},
	)

	w.SetFunction(
		function.Count{
			As: "count",
		},
	)

	go func() {
		for {
			newest := event.Newest(<-w.Output())
			if newest.Int("count") > 10 {
				fmt.Println("Notify!")
			}
		}
	}()

	w.Input() <- LogEvent{
		Time:    time.Now(),
		Level:   1,
		Message: "this is text log.",
	}
}

func LengthWindow() {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := window.NewLength(MyEvent{}, 10)
	defer w.Close()

	w.SetFunction(
		function.AverageInt{
			Name: "Value",
			As:   "avg(Value)",
		},
		function.SumInt{
			Name: "Value",
			As:   "sum(Value)",
		},
	)
}

func View() {
	type MyEvent struct {
		Name  string
		Value int
	}

	w := window.NewTime(MyEvent{}, 10*time.Millisecond)
	defer w.Close()

	w.SetSelector(
		selector.LargerThanInt{
			Name:  "Value",
			Value: 97,
		},
	)
	w.SetFunction(
		function.SelectString{
			Name: "Name",
			As:   "n",
		},
		function.SelectInt{
			Name: "Value",
			As:   "v",
		},
	)
	w.SetView(
		view.OrderByInt{
			Name:    "Value",
			Reverse: true,
		},
		view.Limit{
			Limit:  10,
			Offset: 5,
		},
	)

	go func() {
		for {
			fmt.Println(<-w.Output())
		}
	}()

	for i := 0; i < 100; i++ {
		w.Input() <- MyEvent{
			Name:  "name",
			Value: i,
		}
	}
}

func Builder() {
	b := builder.New()
	b.SetField("Name", reflect.TypeOf(""))
	b.SetField("Value", reflect.TypeOf(0))
	s := b.Build()

	i := s.NewInstance()
	i.SetString("Name", "foobar")
	i.SetInt("Value", 123)

	fmt.Printf("%#v\n", i.Value())
	fmt.Printf("%#v\n", i.Pointer())
}

func Query() {
	type MyEvent struct {
		Name  string
		Value int
	}

	p := parser.New()
	p.Register("MyEvent", MyEvent{})

	query := "select * from MyEvent.length(10)"
	statement, err := p.Parse(query)
	if err != nil {
		log.Println("failed.")
		return
	}

	window := statement.New()
	defer window.Close()
}
