package xyz.lana.challenge.robertom.model;

import lombok.Data;

import java.util.List;

@Data
public class Basket {

    private Long id;
    private List<Item> items;

}
