package xyz.lana.challenge.robertom.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.model.Basket;

@Service
@Slf4j
public class BasketService {

    public Basket create() {
        return null;
    }

    public void addItem(Long basketId, AddItemRequest addItemRequest) {
        //TODO
    }

    public int calculateTotal(long basketId) {
        //TODO
        return 1000;
    }

    public void deleteBasket(long basketId) {
        //TODO
    }

    public Basket get(long basketId) {
        //TODO
        return null;
    }
}
