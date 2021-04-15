package xyz.lana.challenge.robertom.service.discount;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import xyz.lana.challenge.robertom.config.DiscountsConfig;
import xyz.lana.challenge.robertom.model.Item;

import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

@Component
@Slf4j
public class BuyThreeOrMoreGet25OffDiscount extends AbstractDiscount implements Discount {

    @Autowired
    public BuyThreeOrMoreGet25OffDiscount(DiscountsConfig config) {
        super(config);
    }

    @Override
    public int calculate(List<Item> items) {
        if (isNotApplicable()) {
            return 0;
        }

        return calculateDiscount(items);
    }

    @Override
    public String getCode() {
        return DiscountCode.PROMO2.getCode();
    }

    private int calculateDiscount(List<Item> items) {
        AtomicInteger totalDiscount = new AtomicInteger();
        getItemMap(items).forEach((key, value) -> {

            int quantity = value.size();
            if (quantity >= 3) {
                int totalItemDiscount = key.getPrice() * quantity * 25/100;
                log.info("Applying discount for item={} total={}", key, totalItemDiscount);
                totalDiscount.addAndGet(totalItemDiscount);
            }

        });
        return totalDiscount.get();
    }

}
