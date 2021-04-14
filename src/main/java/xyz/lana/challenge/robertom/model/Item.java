package xyz.lana.challenge.robertom.model;

import lombok.Getter;

@Getter
public enum Item {

    PEN("Lana Pen", 500), TSHIRT("Lana T-Shirt", 2000), MUG("Lana Coffee Mug", 750);

    private String name;
    private Integer price;

    Item(String name, Integer price) {
        this.name = name;
        this.price = price;
    }
}
