package xyz.lana.challenge.robertom.service.discount;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
public class DiscountFactory {

    private final List<Discount> discounts;

    @Autowired
    public DiscountFactory(List<Discount> discounts) {
        this.discounts = discounts;
    }

    public List<Discount> getDiscounts() {
        return discounts;
    }

}
