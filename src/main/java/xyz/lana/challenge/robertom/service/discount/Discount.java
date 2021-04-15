package xyz.lana.challenge.robertom.service.discount;

import xyz.lana.challenge.robertom.model.Item;

import java.util.List;

public interface Discount {

    int calculate(List<Item> items);

}
