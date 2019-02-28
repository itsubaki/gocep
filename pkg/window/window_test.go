package window

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/itsubaki/gocep/pkg/event"
	"github.com/itsubaki/gocep/pkg/function"
	"github.com/itsubaki/gocep/pkg/selector"
	"github.com/itsubaki/gocep/pkg/view"
)

func BenchmarkLengthWindowNoFunction128(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: MapEvent{},
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i

		w.Update(MapEvent{m})
	}

}

func BenchmarkLengthWindowSumInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(1)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)

	w.SetFunction(
		function.SumInt{
			Name: "Value",
			As:   "sum(Value)",
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSumInt64(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(64)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	w.SetFunction(
		function.SumInt{
			Name: "Value",
			As:   "sum(Value)",
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSumInt128(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	w.SetFunction(
		function.SumInt{
			Name: "Value",
			As:   "sum(Value)",
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowSumInt256(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(256)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	w.SetFunction(
		function.SumInt{
			Name: "Value",
			As:   "sum(Value)",
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowAverageMap(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: MapEvent{},
		},
	)
	w.SetFunction(
		function.AverageMapInt{
			Name: "Record",
			Key:  "Value",
			As:   "avg(Value)",
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i

		w.Update(MapEvent{m})
	}

}

func BenchmarkLengthWindowAverageInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	w.SetFunction(
		function.AverageInt{
			Name: "Value",
			As:   "avg(Value)",
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowLargerThanMap(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: MapEvent{},
		},
		selector.LargerThanMapInt{
			Name:  "Record",
			Key:   "Value",
			Value: 100,
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i

		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowLargerThanInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
		selector.LargerThanInt{
			Name:  "Value",
			Value: 100,
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowOrderByMap(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: MapEvent{},
		},
	)
	w.SetView(
		view.OrderByMapInt{
			Name:    "Record",
			Key:     "Value",
			Reverse: false,
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i

		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowOrderByInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	w.SetView(
		view.OrderByInt{
			Name:    "Value",
			Reverse: false,
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func BenchmarkLengthWindowOrderByReverseMap(b *testing.B) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: MapEvent{},
		},
	)
	w.SetView(
		view.OrderByMapInt{
			Name:    "Record",
			Key:     "Value",
			Reverse: true,
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		w.Update(MapEvent{m})
	}
}

func BenchmarkLengthWindowOrderByReverseInt(b *testing.B) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(128)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	w.SetView(
		view.OrderByInt{
			Name:    "Value",
			Reverse: true,
		},
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update(IntEvent{"foobar", i})
	}
}

func TestConcurrency(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(2)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
		selector.LargerThanInt{
			Name:  "Value",
			Value: 1,
		},
	)
	w.SetFunction(
		function.Count{
			As: "count",
		},
	)
	w.SetView(
		view.OrderByInt{
			Name:    "Value",
			Reverse: true,
		},
	)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			w.Input() <- IntEvent{"foo", rand.Int()}
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		<-w.Output()
	}
}

func TestLengthWindow(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(2)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
		selector.LargerThanInt{
			Name:  "Value",
			Value: 1,
		},
	)
	w.SetFunction(
		function.Count{
			As: "count",
		},
	)
	w.SetView(
		view.OrderByInt{
			Name:    "Value",
			Reverse: true,
		},
	)

	events := event.List()
	for i := 0; i < 10; i++ {
		events = w.Update(IntEvent{"foo", i})
	}

	if w.Capacity() != 1024 {
		t.Error(w.Capacity())
	}

	var test = []struct {
		index int
		count int
		value int
	}{
		{0, 2, 9},
		{1, 2, 8},
	}

	for _, tt := range test {
		if events[tt.index].Record["count"] != tt.count {
			t.Error(events)
		}
		if events[tt.index].Int("Value") != tt.value {
			t.Error(events)
		}
	}

	if event.Oldest(w.Event()).Record["count"] != 2 {
		t.Error(w.Event())
	}
}

func TestLengthWindowMap(t *testing.T) {
	type MapEvent struct {
		Record map[string]interface{}
	}

	w := NewLength(2)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: MapEvent{},
		},
		selector.LargerThanMapInt{
			Name:  "Record",
			Key:   "Value",
			Value: 1,
		},
	)
	w.SetFunction(
		function.Count{
			As: "count",
		},
		function.AverageMapInt{
			Name: "Record",
			Key:  "Value",
			As:   "avg(Record:Value)",
		},
	)
	w.SetView(
		view.OrderByMapInt{
			Name:    "Record",
			Key:     "Value",
			Reverse: true,
		},
	)

	events := event.List()
	for i := 0; i < 10; i++ {
		m := make(map[string]interface{})
		m["Value"] = i
		events = w.Update(MapEvent{m})
	}

	var test = []struct {
		index int
		count int
		value int
		avg   float64
	}{
		{0, 2, 9, 8.5},
		{1, 2, 8, 8.5},
	}

	for _, tt := range test {
		if events[tt.index].Record["count"] != tt.count {
			t.Error(events)
		}
		if events[tt.index].MapInt("Record", "Value") != tt.value {
			t.Error(events)
		}
		if events[tt.index].Record["avg(Record:Value)"] != tt.avg {
			t.Error(events)
		}
	}
}

func TestLengthWindowListen(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(2)
	defer w.Close()

	w.SetSelector(selector.EqualsType{Accept: IntEvent{}})
	w.Listen("")
}

func TestLengthBatchWindow(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLengthBatch(2)
	defer w.Close()

	w.SetSelector(selector.EqualsType{Accept: IntEvent{}})

	events := event.List()
	for i := 0; i < 10; i++ {
		events = w.Update(IntEvent{"foo", i})
	}

	if events[0].Int("Value") != 8 {
		t.Error(events)
	}

	if events[1].Int("Value") != 9 {
		t.Error(events)
	}
}

func TestTimeWindow0ms(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewTime(0 * time.Millisecond)
	defer w.Close()

	events := event.List()
	for i := 0; i < 10; i++ {
		events = w.Update(IntEvent{"foo", i})
	}

	if len(events) != 0 {
		t.Error(events)
	}
}

func TestTimeWindow10ms(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewTime(1 * time.Millisecond)
	defer w.Close()

	events := event.List()
	for i := 0; i < 10; i++ {
		events = w.Update(IntEvent{"foo", i})
	}

	if len(events) == 0 {
		t.Error(events)
	}
}

func TestTimeBatchWindow10ms(t *testing.T) {
	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewTimeBatch(4 * time.Millisecond)
	defer w.Close()

	for i := 0; i < 10; i++ {
		w.Update(IntEvent{"foo", i})
	}
}

func TestLengthWindowPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	type IntEvent struct {
		Name  string
		Value int
	}

	w := NewLength(10)
	defer w.Close()

	w.SetSelector(
		selector.EqualsType{
			Accept: IntEvent{},
		},
	)
	// IntEvent and Map Function -> panic!!
	w.SetFunction(
		function.AverageMapInt{
			Name: "Record",
			Key:  "Value",
			As:   "avg(Record:Value)",
		},
	)

	events := w.Update(IntEvent{"foobar", 10})
	if len(events) != 0 {
		t.Error(events)
	}
}