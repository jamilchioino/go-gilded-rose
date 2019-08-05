package gildedrose

type Item struct {
	Name            ItemType
	Quality, SellIn int
}

type ItemType string

const (
	Normal        ItemType = "Normal"
	Sulfuras      ItemType = "Sulfuras, Hand of Ragnaros"
	AgedBrie      ItemType = "Aged Brie"
	BackstagePass ItemType = "Backstage Pass"
	Conjured      ItemType = "Conjured"
)

func NewItem(itemType ItemType, quality int, sellIn int) *Item {
	return &Item{itemType, quality, sellIn}
}

func (item *Item) Tick() {
	if item.Name != AgedBrie && item.Name != BackstagePass {
		if item.Quality > 0 {
			if item.Name != Sulfuras {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == BackstagePass {
				if item.SellIn < 11 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}

	if item.Name != Sulfuras {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != AgedBrie {
			if item.Name != BackstagePass {
				if item.Quality > 0 {
					if item.Name != Sulfuras {
						item.Quality = item.Quality - 1
					}
				}
			} else {
				item.Quality = item.Quality - item.Quality
			}
		} else {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
}
