package xyz.lana.challenge.robertom.repository;

import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;

import java.util.List;

public interface BasketStorage {

    Basket create();

    void addItem(Long basketId, Item item);

    int getTotalAmount(Long basketId);

    void delete(Long basketId);

    List<Item> getAllItems(Long basketId);

}