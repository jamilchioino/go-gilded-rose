package gildedrose

type Item struct {
	Name            string
	Quality, SellIn int
}

func NewItem(itemName string, quality int, sellIn int) *Item {
	return &Item{itemName, quality, sellIn}
}

func (item *Item) Tick() {
	if item.Name != "Aged Brie" && item.Name != "Backstage Pass" {
		if item.Quality > 0 {
			if item.Name != "Sulfuras, Hand of Ragnaros" {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.Name == "Backstage Pass" {
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

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage Pass" {
				if item.Quality > 0 {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
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
