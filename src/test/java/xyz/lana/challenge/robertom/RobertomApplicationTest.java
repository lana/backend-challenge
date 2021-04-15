package xyz.lana.challenge.robertom;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import xyz.lana.challenge.robertom.api.BasketController;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
class RobertomApplicationTest {

    @Autowired
    private BasketController controller;

    @Test
    void contextLoads() {
        assertThat(controller).isNotNull();
    }
}