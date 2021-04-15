package xyz.lana.challenge.robertom.api;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.api.response.BasketCreationResponse;
import xyz.lana.challenge.robertom.api.response.BasketTotalAmountResponse;
import xyz.lana.challenge.robertom.converter.CurrencyFormatter;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.service.BasketService;

import javax.validation.Valid;
import java.net.URI;

@RestController
@RequestMapping("/api/basket")
@Slf4j
public class BasketController {

    public static final String BASKET_CREATED_SUCCESSFULLY = "Basket created successfully";
    public static final String BASKET_COULD_NOT_BE_CREATED = "Basket could not be created";

    private final BasketService basketService;
    private final CurrencyFormatter currencyFormatter;

    @Autowired
    public BasketController(BasketService basketService,
                            CurrencyFormatter currencyFormatter) {
        this.basketService = basketService;
        this.currencyFormatter = currencyFormatter;
    }

    @PostMapping
    public ResponseEntity<BasketCreationResponse> create() {
        log.info("Received create basket request");
        Basket basket = basketService.create();
        if (basket == null) {
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(failedBasketResponse());
        }

        URI location = URI.create(String.format("/api/basket/%d", basket.getId()));
        return ResponseEntity.created(location).body(successBasketResponse(basket.getId()));
    }

    @PatchMapping(value = "/{basketId}")
    public ResponseEntity<Void> addItem(@PathVariable Long basketId,
                                        @RequestBody @Valid AddItemRequest addItemRequest) {
        log.info("Received addItem request. basketId={}, addItemRequest={}", basketId, addItemRequest);
        basketService.addItem(basketId, addItemRequest);

        return ResponseEntity.noContent().build();
    }

    @GetMapping(value = "/total/{basketId}")
    public ResponseEntity<BasketTotalAmountResponse> getTotalAmount(@PathVariable Long basketId) {
        log.info("Received getTotalAmount bucket request. basketId={}", basketId);
        int totalAmount = basketService.calculateTotal(basketId);

        String formattedAmount = currencyFormatter.parse(totalAmount);
        BasketTotalAmountResponse response = new BasketTotalAmountResponse();
        response.setTotalAmount(formattedAmount);

        return ResponseEntity.ok(response);
    }

    @DeleteMapping(value = "/{basketId}")
    public ResponseEntity<Void> delete(@PathVariable Long basketId) {
        log.info("Received delete bucket request. basketId={}", basketId);
        basketService.deleteBasket(basketId);

        return ResponseEntity.noContent().build();
    }

    @GetMapping(value = "/{basketId}")
    public ResponseEntity<Basket> get(@PathVariable Long basketId) {
        log.info("Received get bucket request. basketId={}", basketId);
        Basket basket = basketService.get(basketId);
        if (basket == null) {
            return ResponseEntity.notFound().build();
        }

        return ResponseEntity.ok(basket);
    }

    private BasketCreationResponse failedBasketResponse() {
        BasketCreationResponse response = new BasketCreationResponse();
        response.setResponseMsg(BASKET_COULD_NOT_BE_CREATED);

        return response;
    }

    private BasketCreationResponse successBasketResponse(Long id) {
        BasketCreationResponse response = new BasketCreationResponse();
        response.setId(id);
        response.setResponseMsg(BASKET_CREATED_SUCCESSFULLY);

        return response;
    }
}