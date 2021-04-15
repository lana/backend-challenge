package xyz.lana.challenge.robertom.config;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import xyz.lana.challenge.robertom.model.Item;

import java.util.HashSet;
import java.util.Set;

@Getter
@Setter
@ToString
public class DiscountProperties {

    private boolean enable = false;
    private Set<Item> applyTo = new HashSet<>();

}
