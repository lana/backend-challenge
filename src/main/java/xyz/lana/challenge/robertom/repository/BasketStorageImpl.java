package xyz.lana.challenge.robertom.repository;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import xyz.lana.challenge.robertom.exceptions.NotFoundException;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicLong;

@Component
@Slf4j
public class BasketStorageImpl implements BasketStorage {

    public static final String BASKET_ID_COULD_NOT_BE_FOUND = "BasketId %d could not be found";

    private static final AtomicLong counter = new AtomicLong();
    private static final Map<Long, Basket> map = new LinkedHashMap<>();

    public BasketStorageImpl() {
        counter.set(0);
    }

    @Override
    public Basket create() {
        Long newId = counter.incrementAndGet();

        Basket basket = new Basket();
        basket.setId(newId);
        basket.setItems(new ArrayList<>());

        map.put(newId, basket);

        return basket;
    }

    @Override
    public void addItem(Long basketId, Item item) {
        Basket basket = map.get(basketId);
        if (basket == null) {
            throw new NotFoundException(String.format(BASKET_ID_COULD_NOT_BE_FOUND, basketId));
        }
        basket.getItems().add(item);
    }

    @Override
    public void delete(Long basketId) {
        map.remove(basketId);
    }

    @Override
    public void deleteAll() {
        counter.set(0);
        map.clear();
    }

    public Basket get(Long basketId) {
        Basket basket = map.get(basketId);
        if (basket == null) {
            throw new NotFoundException(String.format(BASKET_ID_COULD_NOT_BE_FOUND, basketId));
        }
        return basket;
    }

    @Override
    public List<Item> getAllItems(Long basketId) {
        Basket basket = map.get(basketId);
        if (basket == null) {
            throw new NotFoundException(String.format(BASKET_ID_COULD_NOT_BE_FOUND, basketId));
        }
        return basket.getItems();
    }

}
