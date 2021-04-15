package xyz.lana.challenge.robertom.repository;

import xyz.lana.challenge.robertom.model.Basket;
import xyz.lana.challenge.robertom.model.Item;

import java.util.List;

public interface BasketStorage {

    Basket create();

    void addItem(Long basketId, Item item);

    void delete(Long basketId);

    Basket get(Long basketId);

    List<Item> getAllItems(Long basketId);

}