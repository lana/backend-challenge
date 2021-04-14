package xyz.lana.challenge.robertom.api;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.api.response.BasketResponse;
import xyz.lana.challenge.robertom.converter.CurrencyFormatter;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;
import xyz.lana.challenge.robertom.service.BasketService;

import java.net.URI;
import java.util.Arrays;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.fail;
import static org.mockito.Mockito.*;
import static xyz.lana.challenge.robertom.api.BasketController.*;

@ExtendWith(MockitoExtension.class)
class BasketControllerTest {

    private static final long BASKET_ID = 1L;

    @Mock
    private BasketService basketService;

    @Mock
    private CurrencyFormatter currencyFormatter;

    private BasketController basketController;

    @BeforeEach
    void setUp() {
        basketController = new BasketController(basketService, currencyFormatter);
    }

    @Test
    @DisplayName("When basket is successfully created then status is 201 created")
    void whenBasketIsSuccessfullyCreatedThenStatusIs201Created() {
        when(basketService.create()).thenReturn(basket(BASKET_ID));

        ResponseEntity<BasketResponse> response = basketController.create();

        verify(basketService, times(1)).create();
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.CREATED);
    }

    @Test
    @DisplayName("When basket is successfully created then location is returned")
    void whenBasketIsSuccessfullyCreatedThenLocationIsReturned() {
        when(basketService.create()).thenReturn(basket(BASKET_ID));

        ResponseEntity<BasketResponse> response = basketController.create();

        URI location = URI.create("/api/basket/" + BASKET_ID);
        assertThat(response.getHeaders().getLocation()).isEqualTo(location);
    }

    @Test
    @DisplayName("When basket is created successfully then parameters are the expected")
    void whenBasketIsCreatedSuccessfullyThenParametersAreTheExpected() {
        when(basketService.create()).thenReturn(basket(BASKET_ID));

        ResponseEntity<BasketResponse> response = basketController.create();

        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getId()).isEqualTo(BASKET_ID);
        assertThat(response.getBody().getResponseMsg()).isEqualTo(BASKET_CREATED_SUCCESSFULLY);
    }

    @Test
    @DisplayName("When basket cannot be created then failed response")
    void whenBasketCannotBeCreatedThenFailedResponse() {
        when(basketService.create()).thenReturn(null);

        ResponseEntity<BasketResponse> response = basketController.create();

        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getId()).isNull();
        assertThat(response.getBody().getResponseMsg()).isEqualTo(BASKET_COULD_NOT_BE_CREATED);
    }

    @Test
    @DisplayName("When basket cannot be created then error status")
    void whenBasketCannotBeCreatedThenErrorStatus() {
        when(basketService.create()).thenReturn(null);

        ResponseEntity<BasketResponse> response = basketController.create();

        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @Test
    @DisplayName("When basket creation is called then service is called")
    void whenBasketIsCreatedThenServiceIsCalled() {
        when(basketService.create()).thenReturn(null);

        basketController.create();

        verify(basketService, times(1)).create();
    }

    @Test
    @DisplayName("When adding item then service is called")
    void whenAddingItemThenServiceIsCalled() {
        AddItemRequest addItemRequest = addItemRequest(Item.PEN, 1);

        basketController.addItem(BASKET_ID, addItemRequest);

        verify(basketService, times(1)).addItem(BASKET_ID, addItemRequest);
    }

    @Test
    @DisplayName("When adding item then status is 204 no content")
    void whenAddingItemThenStatusIs204NoContent() {
        AddItemRequest addItemRequest = addItemRequest(Item.PEN, 1);

        ResponseEntity<Void> response = basketController.addItem(BASKET_ID, addItemRequest);

        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.NO_CONTENT);
    }

    @Test
    @DisplayName("When getting basket total amount then service is called")
    void whenGettingBasketTotalAmountThenServiceIsCalled() {
        basketController.getTotalAmount(BASKET_ID);

        verify(basketService, times(1)).calculateTotal(BASKET_ID);
    }

    @Test
    @DisplayName("When getting basket total amount then currency formatter is called")
    void whenGettingBasketTotalAmountThenCurrencyFormatterIsCalled() {
        when(basketService.calculateTotal(BASKET_ID)).thenReturn(1000);

        basketController.getTotalAmount(BASKET_ID);

        verify(currencyFormatter, times(1)).parse(1000);
    }

    @Test
    @DisplayName("When getting basket total amount then status is 200 Ok")
    void whenGettingBasketTotalAmountThenStatusIs200Ok() {
        when(basketService.calculateTotal(BASKET_ID)).thenReturn(1000);

        ResponseEntity<String> response = basketController.getTotalAmount(BASKET_ID);

        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
    }

    @Test
    @DisplayName("When basket is removed then service is called")
    void whenBasketIsRemovedThenServiceIsCalled() {
        basketController.delete(BASKET_ID);

        verify(basketService, times(1)).deleteBasket(BASKET_ID);
    }

    @Test
    @DisplayName("When basket is removed then status is 200 ok")
    void whenBasketIsRemovedThenStatusIs200Ok() {
        ResponseEntity<Void> response = basketController.delete(BASKET_ID);

        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.NO_CONTENT);
    }

    @Test
    @DisplayName("When getting basket then service is called")
    void whenGettingBasketThenServiceIsCalled() {
        basketController.get(BASKET_ID);

        verify(basketService, times(1)).get(BASKET_ID);
    }

    @Test
    @DisplayName("When getting existing basket then status is 200 ok")
    void whenGettingExistingBasketThenStatusIs200Ok() {
        when(basketService.get(BASKET_ID)).thenReturn(basket(BASKET_ID));
        
        ResponseEntity<Basket> response = basketController.get(BASKET_ID);
        
        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
    }

    @Test
    @DisplayName("When getting non existing basket then status 404 not found is returned")
    void whenGettingNonExistingBasketThenStatus404NotFoundIsReturned() {
        when(basketService.get(BASKET_ID)).thenReturn(null);

        ResponseEntity<Basket> response = basketController.get(BASKET_ID);

        assertThat(response.getStatusCode()).isEqualTo(HttpStatus.NOT_FOUND);
    }

    @Test
    @DisplayName("When getting existing basket then basket is returned")
    void whenGettingExistingBasketThenBasketIsReturned() {
        Basket basket = basket(BASKET_ID, Item.PEN);

        when(basketService.get(BASKET_ID)).thenReturn(basket);

        ResponseEntity<Basket> response = basketController.get(BASKET_ID);
        
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getId()).isEqualTo(BASKET_ID);
        assertThat(response.getBody().getItems()).containsExactly(Item.PEN);

    }

    private AddItemRequest addItemRequest(Item code, int quantity) {
        AddItemRequest addItemRequest = new AddItemRequest();
        addItemRequest.setCode(code);
        addItemRequest.setQuantity(quantity);

        return addItemRequest;
    }

    private Basket basket(Long id, Item... items){
        Basket basket = new Basket();
        basket.setId(id);
        basket.setItems(Arrays.asList(items));

        return basket;
    }
}