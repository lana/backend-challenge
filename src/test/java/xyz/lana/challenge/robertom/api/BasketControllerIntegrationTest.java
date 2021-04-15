package xyz.lana.challenge.robertom.api;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import xyz.lana.challenge.robertom.RobertomApplication;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.api.resource.ErrorResource;
import xyz.lana.challenge.robertom.api.response.BasketCreationResponse;
import xyz.lana.challenge.robertom.api.response.BasketTotalAmountResponse;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;
import xyz.lana.challenge.robertom.repository.BasketStorage;

import java.net.URI;

import static org.assertj.core.api.Assertions.assertThat;
import static xyz.lana.challenge.robertom.api.BasketController.BASKET_CREATED_SUCCESSFULLY;
import static xyz.lana.challenge.robertom.repository.BasketStorageImpl.BASKET_ID_COULD_NOT_BE_FOUND;

@SpringBootTest(classes = RobertomApplication.class,
        webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class BasketControllerIntegrationTest {

    private static final String BASE_PATH = "/api/basket";
    private static final String TOTAL_PATH = BASE_PATH + "/total";
    private static final long BASKET_ID = 1L;

    @LocalServerPort
    private int port;

    @Autowired
    private TestRestTemplate restTemplate;

    @Autowired
    private BasketStorage basketStorage;

    @BeforeEach
    public void setup() {
        restTemplate.getRestTemplate().setRequestFactory(new HttpComponentsClientHttpRequestFactory());
    }

    @AfterEach
    public void cleanUpEach(){
        basketStorage.deleteAll();
    }

    @Test
    @DisplayName("Create a new basket")
    void createANewBasket() {
        String url = "http://localhost:" + port + BASE_PATH;

        ResponseEntity<BasketCreationResponse> response = restTemplate.exchange(url,
                HttpMethod.POST,
                null,
                new ParameterizedTypeReference<BasketCreationResponse>() {});

        assertThat(response.getHeaders().getLocation()).isEqualTo(URI.create("/api/basket/1"));
        assertThat(response.getStatusCodeValue()).isEqualTo(201);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getId()).isOne();
        assertThat(response.getBody().getResponseMsg()).isEqualTo(BASKET_CREATED_SUCCESSFULLY);
    }

    @Test
    @DisplayName("Adding item for an existing basket")
    void addingItemForAnExistingBasket() {
        Basket basket = basketStorage.create();
        AddItemRequest addItemRequest = addItemRequest();
        String url = "http://localhost:" + port + BASE_PATH + "/" + basket.getId();

        ResponseEntity<Void> response = restTemplate.exchange(url,
                HttpMethod.PATCH,
                new HttpEntity<>(addItemRequest),
                new ParameterizedTypeReference<Void>() {});

        assertThat(response.getStatusCodeValue()).isEqualTo(204);
        assertThat(response.getBody()).isNull();
    }

    @Test
    @DisplayName("Get basket total amount")
    void getBasketTotalAmount() {
        Basket basket = basketStorage.create();
        basketStorage.addItem(basket.getId(), Item.PEN);
        basketStorage.addItem(basket.getId(), Item.TSHIRT);
        basketStorage.addItem(basket.getId(), Item.PEN);
        basketStorage.addItem(basket.getId(), Item.PEN);
        basketStorage.addItem(basket.getId(), Item.MUG);
        basketStorage.addItem(basket.getId(), Item.TSHIRT);
        basketStorage.addItem(basket.getId(), Item.TSHIRT);

        String url = "http://localhost:" + port + TOTAL_PATH + "/" + basket.getId();

        ResponseEntity<BasketTotalAmountResponse> response = restTemplate.exchange(url,
                HttpMethod.GET,
                null,
                new ParameterizedTypeReference<BasketTotalAmountResponse>() {});

        assertThat(response.getStatusCodeValue()).isEqualTo(200);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getTotalAmount()).isEqualTo("62.50€");
    }

    @Test
    @DisplayName("Delete basket")
    void deleteBasket() {
        Basket basket = basketStorage.create();

        String url = "http://localhost:" + port + BASE_PATH + "/" + basket.getId();

        ResponseEntity<Void> response = restTemplate.exchange(url,
                HttpMethod.DELETE,
                null,
                new ParameterizedTypeReference<Void>() {});

        assertThat(response.getStatusCodeValue()).isEqualTo(204);
        assertThat(response.getBody()).isNull();
    }

    @Test
    @DisplayName("Get existing basket")
    void getExistingBasket() {
        Basket basket = basketStorage.create();

        System.out.println(basket);

        basketStorage.addItem(basket.getId(), Item.PEN);

        String url = "http://localhost:" + port + BASE_PATH + "/" + basket.getId();

        ResponseEntity<Basket> response = restTemplate.exchange(url,
                HttpMethod.GET,
                null,
                new ParameterizedTypeReference<Basket>() {});

        assertThat(response.getStatusCodeValue()).isEqualTo(200);
        assertThat(response.getBody()).isNotNull();
        assertThat(response.getBody().getId()).isOne();
        assertThat(response.getBody().getItems()).hasSize(1);
        assertThat(response.getBody().getItems()).containsExactly(Item.PEN);
    }

    @Test
    @DisplayName("Get non existing basket")
    void getNonExistingBasket() {
        String url = "http://localhost:" + port + BASE_PATH + "/" + BASKET_ID;

        ResponseEntity<ErrorResource> response = restTemplate.exchange(url,
                HttpMethod.GET,
                null,
                new ParameterizedTypeReference<ErrorResource>() {});

        assertThat(response.getStatusCodeValue()).isEqualTo(404);
        assertThat(response.getBody()).isNotNull();

        String message = String.format(BASKET_ID_COULD_NOT_BE_FOUND, BASKET_ID);
        assertThat(response.getBody().getMessage()).isEqualTo(message);
    }

    private AddItemRequest addItemRequest() {
        AddItemRequest addItemRequest = new AddItemRequest();
        addItemRequest.setCode(Item.PEN.name());
        addItemRequest.setQuantity(1);

        return addItemRequest;
    }

}