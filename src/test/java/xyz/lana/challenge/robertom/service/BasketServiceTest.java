package xyz.lana.challenge.robertom.service;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.config.DiscountProperties;
import xyz.lana.challenge.robertom.config.DiscountsConfig;
import xyz.lana.challenge.robertom.exceptions.NotFoundException;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;
import xyz.lana.challenge.robertom.repository.BasketStorage;
import xyz.lana.challenge.robertom.service.discount.BuyThreeOrMoreGet25OffDiscount;
import xyz.lana.challenge.robertom.service.discount.BuyTwoGetOneFreeDiscount;
import xyz.lana.challenge.robertom.service.discount.DiscountCode;
import xyz.lana.challenge.robertom.service.discount.DiscountFactory;

import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class BasketServiceTest {

    private static final long BASKET_ID = 1L;
    private static final String NOT_FOUND_ERROR_MESSAGE = "not-found-error-message";

    @Mock
    private BasketStorage basketStorage;

    @Mock
    private DiscountFactory discountFactory;

    private BasketService basketService;

    @BeforeEach
    void setUp() {
        basketService = new BasketService(basketStorage, discountFactory);
    }

    @Test
    @DisplayName("When create basket then repository is called")
    void whenCreateBasketThenRepositoryIsCalled() {
        basketService.create();

        verify(basketStorage, times(1)).create();
    }


    @Test
    @DisplayName("When create basket successfully then basket is returned with the expected parameters")
    void whenCreateBasketSuccessfullyThenBasketIsReturnedWithTheExpectedParameters() {
        when(basketStorage.create()).thenReturn(basket());

        Basket basket = basketService.create();

        assertThat(basket).isNotNull();
        assertThat(basket.getId()).isEqualTo(BASKET_ID);
        assertThat(basket.getItems()).isEmpty();
    }

    @Test
    @DisplayName("When basket cannot be created then null is returned")
    void whenBasketCannotBeCreatedThenNullIsReturned() {
        when(basketStorage.create()).thenReturn(null);

        Basket basket = basketService.create();

        assertThat(basket).isNull();
    }

    @Test
    @DisplayName("When add new item then repository is called")
    void whenAddNewItemThenRepositoryIsCalled() {
        AddItemRequest addItemRequest = addItemRequest();
        basketService.addItem(BASKET_ID, addItemRequest);

        verify(basketStorage, times(2)).addItem(BASKET_ID, Item.PEN);
    }

    @Test
    @DisplayName("When adding new item for a non existing basket then not found exception")
    void whenAddingNewItemForANonExistingBasketThenNotFoundException() {
        AddItemRequest addItemRequest = addItemRequest();

        doThrow(new NotFoundException(NOT_FOUND_ERROR_MESSAGE)).when(basketStorage).addItem(BASKET_ID, Item.PEN);

        Throwable thrown = catchThrowable(() -> basketService.addItem(BASKET_ID, addItemRequest));

        assertThat(thrown).isExactlyInstanceOf(NotFoundException.class);
        assertThat(thrown.getMessage()).isEqualTo(NOT_FOUND_ERROR_MESSAGE);
    }

    @Test
    @DisplayName("When delete basket then repository is called")
    void whenDeleteBasketThenRepositoryIsCalled() {
        basketService.deleteBasket(BASKET_ID);

        verify(basketStorage, times(1)).delete(BASKET_ID);
    }

    @Test
    @DisplayName("When getting basket then repository is called")
    void whenGettingBasketThenRepositoryIsCalled() {
        basketService.get(BASKET_ID);

        verify(basketStorage, times(1)).get(BASKET_ID);
    }

    @Test
    @DisplayName("When getting existing basket then it's returned")
    void whenGettingExistingBasketThenItSReturned() {
        when(basketStorage.get(BASKET_ID)).thenReturn(basket(Item.PEN));

        Basket basket = basketService.get(BASKET_ID);

        assertThat(basket).isNotNull();
        assertThat(basket.getId()).isEqualTo(BASKET_ID);
        assertThat(basket.getItems()).containsExactly(Item.PEN);
    }

    @Test
    @DisplayName("When getting a non existing basket then it's not returned")
    void whenGettingANonExistingBasketThenItSNotReturned() {
        doThrow(new NotFoundException(NOT_FOUND_ERROR_MESSAGE)).when(basketStorage).get(BASKET_ID);

        Throwable thrown = catchThrowable(() -> basketService.get(BASKET_ID));

        assertThat(thrown).isExactlyInstanceOf(NotFoundException.class);
        assertThat(thrown.getMessage()).isEqualTo(NOT_FOUND_ERROR_MESSAGE);
    }

    @Test
    @DisplayName("When calculate total amount then services is called")
    void whenCalculateTotalAmountThenServicesIsCalled() {
        when(basketStorage.get(BASKET_ID)).thenReturn(basket(Item.PEN));

        basketService.calculateTotal(BASKET_ID);

        verify(basketStorage, times(1)).get(BASKET_ID);
    }

    @Test
    @DisplayName("When calculate total without discounts then total is the sum of the prices for all the basket items")
    void whenCalculateTotalWithoutDiscountsThenTotalIsTheSumOfThePricesForAllTheBasketItems() {
        when(basketStorage.get(BASKET_ID)).thenReturn(basket(Item.PEN, Item.MUG, Item.TSHIRT));

        int total = basketService.calculateTotal(BASKET_ID);

        assertThat(total).isEqualTo(Item.PEN.getPrice() + Item.MUG.getPrice() + Item.TSHIRT.getPrice());
    }

    @Test
    @DisplayName("When calculate total using buy 2 get 1 free discount then it's applied")
    void whenCalculateTotalUsingBuy2Get1FreeDiscountThenItSApplied() {
        DiscountsConfig discountsConfig = discountsConfig();

        when(basketStorage.get(BASKET_ID)).thenReturn(basket(Item.PEN, Item.PEN, Item.PEN));
        when(discountFactory.getDiscounts()).thenReturn(Collections.singletonList(new BuyTwoGetOneFreeDiscount(discountsConfig)));

        int total = basketService.calculateTotal(BASKET_ID);

        assertThat(total).isEqualTo(Item.PEN.getPrice()*2);
    }

    @Test
    @DisplayName("When calculate total using buy three or more get 25% off then it's applied ")
    void whenCalculateTotalUsingBuyThreeOrMoreGet25OffThenItSApplied() {
        DiscountsConfig discountsConfig = discountsConfig();

        when(basketStorage.get(BASKET_ID)).thenReturn(basket(Item.TSHIRT, Item.TSHIRT, Item.TSHIRT));
        when(discountFactory.getDiscounts()).thenReturn(Collections.singletonList(new BuyThreeOrMoreGet25OffDiscount(discountsConfig)));

        int total = basketService.calculateTotal(BASKET_ID);

        int expectedPrice = (Item.TSHIRT.getPrice() + Item.TSHIRT.getPrice() + Item.TSHIRT.getPrice()) -
                (Item.TSHIRT.getPrice() + Item.TSHIRT.getPrice() + Item.TSHIRT.getPrice()) * 25/100;
        assertThat(total).isEqualTo(expectedPrice);
    }

    @Test
    @DisplayName("When calculate total with no discounts then it's calculated without any discount")
    void whenCalculateTotalWithNoDiscountsThenItSCalculatedWithoutAnyDiscount() {
        DiscountsConfig discountsConfig = discountsConfig();
        discountsConfig.getDiscounts().get(DiscountCode.PROMO1.getCode()).setEnable(false);
        discountsConfig.getDiscounts().get(DiscountCode.PROMO2.getCode()).setEnable(false);

        when(basketStorage.get(BASKET_ID)).thenReturn(basket(Item.PEN, Item.PEN, Item.TSHIRT, Item.TSHIRT, Item.TSHIRT));
        when(discountFactory.getDiscounts()).thenReturn(Collections.singletonList(new BuyTwoGetOneFreeDiscount(discountsConfig)));

        int total = basketService.calculateTotal(BASKET_ID);

        int expectedPrice = Item.PEN.getPrice() +
                Item.PEN.getPrice() +
                Item.TSHIRT.getPrice() +
                Item.TSHIRT.getPrice() +
                Item.TSHIRT.getPrice();
        assertThat(total).isEqualTo(expectedPrice);
    }

    private Basket basket(Item... items){
        Basket basket = new Basket();
        basket.setId(BasketServiceTest.BASKET_ID);
        basket.setItems(Arrays.asList(items));

        return basket;
    }

    private AddItemRequest addItemRequest() {
        AddItemRequest addItemRequest = new AddItemRequest();
        addItemRequest.setCode(Item.PEN.name());
        addItemRequest.setQuantity(2);

        return addItemRequest;
    }

    private DiscountsConfig discountsConfig(){
        Map<String, DiscountProperties> map = new HashMap<>();
        map.put(DiscountCode.PROMO1.getCode(), discountProperties(new HashSet<>(Collections.singletonList(Item.PEN))));
        map.put(DiscountCode.PROMO2.getCode(), discountProperties(new HashSet<>(Collections.singletonList(Item.TSHIRT))));

        DiscountsConfig discountsConfig = new DiscountsConfig();
        discountsConfig.setDiscounts(map);

        return discountsConfig;
    }

    private DiscountProperties discountProperties(Set<Item> applyTo){
        DiscountProperties discountProperties = new DiscountProperties();
        discountProperties.setEnable(true);
        discountProperties.setApplyTo(applyTo);

        return discountProperties;
    }

}