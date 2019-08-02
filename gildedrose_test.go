package gildedrose_test

import (
	"testing"

	gildedrose "github.com/jamilchioino/go-gilded-rose"

	"github.com/matryer/is"
)

var itemTests = []struct {
	testName   string
	itemType   gildedrose.ItemType
	inQuality  int
	inSellIn   int
	outQuality int
	outSellIn  int
}{
	{"updates normal items before sell date", gildedrose.Normal, 10, 5, 9, 4},
	{"updates normal items on the sell date", gildedrose.Normal, 10, 0, 8, -1},
	{"updates normal items after the sell date", gildedrose.Normal, 10, -5, 8, -6},
	{"updates normal items with a quality of 0", gildedrose.Normal, 0, 5, 0, 4},

	{"updates Brie items before the sell date", gildedrose.AgedBrie, 10, 5, 11, 4},
	{"updates Brie items before the sell date with maximum quality", gildedrose.AgedBrie, 10, 0, 12, -1},
	{"updates Brie items on the sell date, near maximum quality", gildedrose.AgedBrie, 49, 0, 50, -1},
	{"updates Brie items on the sell date with maximum quality", gildedrose.AgedBrie, 50, 0, 50, -1},
	{"updates Brie items after the sell date", gildedrose.AgedBrie, 10, -10, 12, -11},
	{"updates Brie items after the sell date with maximum quality", gildedrose.AgedBrie, 50, -10, 50, -11},

	{"updates Sulfuras items before the sell date", gildedrose.Sulfuras, 10, 5, 10, 5},
	{"updates Sulfuras items on the sell date", gildedrose.Sulfuras, 10, 0, 10, 0},
	{"updates Sulfuras items after the sell date", gildedrose.Sulfuras, 10, -1, 10, -1},

	{"updates Backstage pass items long before the sell date", gildedrose.BackstagePass, 10, 11, 11, 10},
	{"updates Backstage pass items close to the sell date", gildedrose.BackstagePass, 10, 10, 12, 9},
	{"updates Backstage pass items close to the sell data, at max quality", gildedrose.BackstagePass, 50, 10, 50, 9},
	{"updates Backstage pass items very close to the sell date", gildedrose.BackstagePass, 10, 5, 13, 4},
	{"updates Backstage pass items very close to the sell date, at max quality", gildedrose.BackstagePass, 50, 5, 50, 4},
	{"updates Backstage pass items with one day left to sell", gildedrose.BackstagePass, 10, 1, 13, 0},
	{"updates Backstage pass items with one day left to sell, at max quality", gildedrose.BackstagePass, 50, 1, 50, 0},
	{"updates Backstage pass items on the sell date", gildedrose.BackstagePass, 10, 0, 0, -1},
	{"updates Backstage pass items after the sell date", gildedrose.BackstagePass, 10, -1, 0, -2},

	// {"updates Conjured items before the sell date", gildedrose.Conjured, 10, 10, 8, 9},
	// {"updates Conjured items at zero quality", gildedrose.Conjured, 0, 10, 0, 9},
	// {"updates Conjured items on the sell date", gildedrose.Conjured, 10, 0, 6, -1},
	// {"updates Conjured items on the sell date at 0 quality", gildedrose.Conjured, 0, 0, 0, -1},
	// {"updates Conjured items after the sell date", gildedrose.Conjured, 10, -10, 6, -11},
	// {"updates Conjured items after the sell date at zero quality", gildedrose.Conjured, 0, -10, 0, -11},
}

func TestItems(t *testing.T) {
	for _, tt := range itemTests {
		t.Run(tt.testName, func(t *testing.T) {
			is := is.New(t)
			item := gildedrose.NewItem(tt.itemType, tt.inQuality, tt.inSellIn)
			item.Tick()
			is.Equal(item.Quality, tt.outQuality)
			is.Equal(item.SellIn, tt.outSellIn)
		})
	}
}
