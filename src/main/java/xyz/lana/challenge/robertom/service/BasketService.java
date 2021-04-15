package xyz.lana.challenge.robertom.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.collections4.CollectionUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import xyz.lana.challenge.robertom.api.request.AddItemRequest;
import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;
import xyz.lana.challenge.robertom.repository.BasketStorage;
import xyz.lana.challenge.robertom.service.discount.DiscountFactory;

@Service
@Slf4j
public class BasketService {

    private final BasketStorage basketStorage;
    private final DiscountFactory discountFactory;

    @Autowired
    public BasketService(BasketStorage basketStorage,
                         DiscountFactory discountFactory) {
        this.basketStorage = basketStorage;
        this.discountFactory = discountFactory;
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
            Item itemToAdd = Item.fromText(addItemRequest.getCode());
            basketStorage.addItem(basketId, itemToAdd);
        }
        log.info("Added to the basketId={} the items requested addItemRequest={}", basketId, addItemRequest);
    }

    public void deleteBasket(Long basketId) {
        basketStorage.delete(basketId);
        log.info("BasketId={} deleted", basketId);

    }

    public Basket get(Long basketId) {
        Basket basket = basketStorage.get(basketId);
        log.info("BasketId={} found. basket={}", basketId, basket);
        return basket;
    }

    public int calculateTotal(Long basketId) {
        Basket basket = get(basketId);
        int totalAmountWithoutDiscounts = getTotalAmountWithoutDiscounts(basket);

        if (CollectionUtils.isEmpty(discountFactory.getDiscounts())) {
            return totalAmountWithoutDiscounts;
        }

        int discounts = getDiscounts(basket);
        log.info("Total basket amount calculated. totalAmountWithoutDiscounts={}, discounts={}, final={}",
                totalAmountWithoutDiscounts, discounts, totalAmountWithoutDiscounts - discounts);

        return totalAmountWithoutDiscounts - discounts;
    }

    private Integer getDiscounts(Basket basket) {
        return discountFactory.getDiscounts()
                .stream()
                .map(discount -> discount.calculate(basket.getItems()))
                .reduce(0, Integer::sum);
    }

    private int getTotalAmountWithoutDiscounts(Basket basket) {
        return basket.getItems()
                .stream()
                .reduce(0, (subtotal, item) -> subtotal + item.getPrice(), Integer::sum);
    }
}
