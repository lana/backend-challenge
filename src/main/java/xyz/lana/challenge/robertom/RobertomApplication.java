package xyz.lana.challenge.robertom;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import xyz.lana.challenge.robertom.config.DiscountsConfig;

@SpringBootApplication
@EnableConfigurationProperties(DiscountsConfig.class)
public class RobertomApplication {

    public static void main(String[] args) {
        SpringApplication.run(RobertomApplication.class, args);
    }

}
