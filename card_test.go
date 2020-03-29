package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Seven, Suit: Spade})
	fmt.Println(Card{Rank: King, Suit: Heart})
	fmt.Println(Card{Rank: Queen, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Ace of Hearts
	//Seven of Spades
	//King of Hearts
	//Queen of Hearts
	//Joker
}

func TestNew(t *testing.T){
	cards := New()
	if len(cards) != 13*4{
		t.Error("Wrong number of cards in the deck")
	}
}

func TestDefaultSort(t *testing.T){
	cards := New(DefaultSort)
	exp := Card{Rank:Ace, Suit:Spade}
	if cards[0] != exp{
		t.Error("Expected Ace of Spades. Got ", cards[0])
	}
}