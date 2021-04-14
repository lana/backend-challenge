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
import static org.junit.jupiter.api.Assertions.fail;

class BasketStorageImplTest {

    private static final long NON_EXISTENT_BASKET_ID = -999L;
    private static final Item CODE = Item.PEN;
    private static final String NAME = "some-name";
    private static final int PRICE_1 = 1000;
    private static final int PRICE_2 = 2000;
    private static final int PRICE_ADDED = PRICE_1 + PRICE_2;

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
    @DisplayName("When getting the total amount for a non existent basket then error")
    void whenGettingTheTotalAmountForANonExistentBasketThenError() {
        Throwable thrown = catchThrowable(() -> basketStorage.getTotalAmount(NON_EXISTENT_BASKET_ID));

        assertThat(thrown).isExactlyInstanceOf(NotFoundException.class);
    }

    @Test
    @DisplayName("When getting the total amount then all prices from the same basket are added together")
    void whenGettingTheTotalAmountThenAllPricesFromTheSameBasketAreAddedTogether() {
        Basket basket = basketStorage.create();

        basketStorage.addItem(basket.getId(), Item.PEN);
        basketStorage.addItem(basket.getId(), Item.MUG);

        int total = basketStorage.getTotalAmount(basket.getId());

        assertThat(total).isEqualTo(Item.PEN.getPrice() + Item.MUG.getPrice());
    }

    @Test
    @DisplayName("When deleting basket then it's deleted from the storage")
    void whenDeletingBasketThenItSDeletedFromTheStorage() {
        Basket basket = basketStorage.create();

        basketStorage.delete(basket.getId());

        Throwable thrown = catchThrowable(() -> basketStorage.getAllItems(basket.getId()));

        assertThat(thrown).isExactlyInstanceOf(NotFoundException.class);
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