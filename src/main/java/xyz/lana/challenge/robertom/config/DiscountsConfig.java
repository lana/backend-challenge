package xyz.lana.challenge.robertom.config;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import org.springframework.boot.context.properties.ConfigurationProperties;

import java.util.HashMap;
import java.util.Map;

@ConfigurationProperties(prefix = "basket")
@Getter
@Setter
@ToString
public class DiscountsConfig {

    private Map<String, DiscountProperties> discounts = new HashMap<>();

}
