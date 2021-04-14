package xyz.lana.challenge.robertom.service;

import com.sun.org.apache.bcel.internal.classfile.Code;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.exceptions.NotFoundException;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;
import xyz.lana.challenge.robertom.repository.BasketStorage;

import java.util.Arrays;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.catchThrowable;
import static org.junit.jupiter.api.Assertions.fail;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class BasketServiceTest {

    private static final long BASKET_ID = 1L;
    private static final long NON_EXISTING_BASKET_ID = -999L;
    private static final String NOT_FOUND_ERROR_MESSAGE = "not-found-error-message";

    @Mock
    private BasketStorage basketStorage;

    private BasketService basketService;

    @BeforeEach
    void setUp() {
        basketService = new BasketService(basketStorage);
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
        when(basketStorage.create()).thenReturn(basket(BASKET_ID));

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
        AddItemRequest addItemRequest = addItemRequest(Item.PEN, 2);
        basketService.addItem(BASKET_ID, addItemRequest);

        verify(basketStorage, times(2)).addItem(BASKET_ID, Item.PEN);
    }

    @Test
    @DisplayName("When adding new item for a non existing basket then not found exception")
    void whenAddingNewItemForANonExistingBasketThenNotFoundException() {
        AddItemRequest addItemRequest = addItemRequest(Item.PEN, 2);

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
        when(basketStorage.get(BASKET_ID)).thenReturn(basket(BASKET_ID, Item.PEN));

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

    private Basket basket(Long id, Item... items){
        Basket basket = new Basket();
        basket.setId(id);
        basket.setItems(Arrays.asList(items));

        return basket;
    }

    private AddItemRequest addItemRequest(Item code, int quantity) {
        AddItemRequest addItemRequest = new AddItemRequest();
        addItemRequest.setCode(code);
        addItemRequest.setQuantity(quantity);

        return addItemRequest;
    }

}