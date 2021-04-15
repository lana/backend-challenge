package xyz.lana.challenge.robertom.repository;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import xyz.lana.challenge.robertom.exceptions.NotFoundException;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;

class BasketStorageImplTest {

    private static final long NON_EXISTENT_BASKET_ID = -999L;

    private BasketStorage basketStorage;

    @BeforeEach
    void setUp() {
        basketStorage = new BasketStorageImpl();
    }

    @Test
    @DisplayName("When creating new basket then it's added to the storage")
    void whenCreatingNewBasketThenItSAddedToTheStorage() {
        Basket basket = basketStorage.create();

        assertThat(basket.getId()).isOne();
        assertThat(basket.getItems()).isEmpty();
    }

    @Test
    @DisplayName("When adding a new item to an existent basket then it's added")
    void whenAddingANewItemToAnExistentBasketThenItSAdded() {
        Basket basket = basketStorage.create();

        basketStorage.addItem(basket.getId(), Item.PEN);

        List<Item> basketItems = basketStorage.getAllItems(basket.getId());

        assertThat(basketItems).containsExactly(Item.PEN);
    }

    @Test
    @DisplayName("When adding a new item to an non existent basket then it's not added")
    void whenAddingANewItemToAnNonExistentBasketThenItSNotAdded() {
        Throwable thrown = catchThrowable(() -> basketStorage.addItem(NON_EXISTENT_BASKET_ID, Item.PEN));

        assertThat(thrown).isExactlyInstanceOf(NotFoundException.class);
    }

    @Test
    @DisplayName("When deleting basket then it's deleted from the storage")
    void whenDeletingBasketThenItSDeletedFromTheStorage() {
        Basket basket = basketStorage.create();

        basketStorage.delete(basket.getId());

        Throwable thrown = catchThrowable(() -> basketStorage.get(basket.getId()));

        assertThat(thrown).isExactlyInstanceOf(NotFoundException.class);
    }

    @Test
    @DisplayName("When deleting all baskets then they're deleted from storage")
    void whenDeletingAllBasketsThenTheyReDeletedFromStorage() {
        Basket basket1 = basketStorage.create();
        Basket basket2 = basketStorage.create();

        basketStorage.deleteAll();

        Throwable thrown1 = catchThrowable(() -> basketStorage.get(basket1.getId()));
        assertThat(thrown1).isExactlyInstanceOf(NotFoundException.class);

        Throwable thrown2 = catchThrowable(() -> basketStorage.get(basket2.getId()));
        assertThat(thrown2).isExactlyInstanceOf(NotFoundException.class);
    }

    @Test
    @DisplayName("When getting all the items in the basket then they are returned in a list")
    void whenGettingAllTheItemsInTheBasketThenTheyAreReturnedInAList() {

        Basket basket = basketStorage.create();

        basketStorage.addItem(basket.getId(), Item.PEN);
        basketStorage.addItem(basket.getId(), Item.MUG);

        List<Item> items = basketStorage.getAllItems(basket.getId());

        assertThat(items).containsExactly(Item.PEN, Item.MUG);
    }
}