package xyz.lana.challenge.robertom.service.discount;

import xyz.lana.challenge.robertom.config.DiscountProperties;
import xyz.lana.challenge.robertom.config.DiscountsConfig;
import xyz.lana.challenge.robertom.model.Item;

import java.util.List;
import java.util.Map;
import java.util.function.Function;

import static java.util.stream.Collectors.groupingBy;

public abstract class AbstractDiscount {

    private final DiscountsConfig config;

    AbstractDiscount(DiscountsConfig config) {
        this.config = config;
    }

    public abstract String getCode();

    protected DiscountProperties getProperties(){
        return config.getDiscounts().get(getCode());
    }

    protected boolean isNotApplicable(){
        DiscountProperties properties = config.getDiscounts().get(getCode());
        return properties == null || !properties.isEnable();
    }

    protected Map<Item, List<Item>> getItemMap(List<Item> items) {
        return items.stream()
                .filter(item -> getProperties().getApplyTo().contains(item))
                .collect(groupingBy(Function.identity()));
    }

}
