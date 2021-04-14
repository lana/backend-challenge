package xyz.lana.challenge.robertom.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.repository.BasketStorage;

@Service
@Slf4j
public class BasketService {

    private final BasketStorage basketStorage;

    @Autowired
    public BasketService(BasketStorage basketStorage) {
        this.basketStorage = basketStorage;
    }

    public Basket create() {
        Basket basket = basketStorage.create();
        if (basket == null) {
            log.error("Basket could not be created");
            return null;
        }

        log.info("Basket={} created successfully", basket);
        return basket;
    }

    public void addItem(Long basketId, AddItemRequest addItemRequest) {
        if (addItemRequest.getQuantity() < 1) {
            log.info("No items for adding to the basketId={}", basketId);
            return;
        }
        for (int i = 0; i < addItemRequest.getQuantity(); i++) {
            basketStorage.addItem(basketId, addItemRequest.getCode());
        }
        log.info("Added to the basketId={} the items requested addItemRequest={}", basketId, addItemRequest);
    }

    public void deleteBasket(long basketId) {
        basketStorage.delete(basketId);
        log.info("BasketId={} deleted", basketId);

    }

    public Basket get(long basketId) {
        Basket basket = basketStorage.get(basketId);
        log.info("BasketId={} found. basket={}", basketId, basket);
        return basket;
    }

    public int calculateTotal(long basketId) {
        //TODO
        return 1000;
    }
}
