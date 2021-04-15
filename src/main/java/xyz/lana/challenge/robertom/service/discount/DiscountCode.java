package xyz.lana.challenge.robertom.service.discount;

import lombok.Getter;

@Getter
public enum DiscountCode {

    PROMO1("buy-two-get-one-free"), PROMO2("buy-three-or-more-get-twenty-five-percent-discount");

    private final String code;

    DiscountCode(String code) {
        this.code = code;
    }
}
