package xyz.lana.challenge.robertom.converter;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;

class CurrencyFormatterTest {

    private static final int MONEY_INTEGER = 10000;
    private static final String MONEY_FORMATTED = "100.00€";

    private CurrencyFormatter currencyFormatter;

    @BeforeEach
    void setUp() {
        currencyFormatter = new CurrencyFormatter();
    }

    @Test
    @DisplayName("When parsin then expected format is returned")
    void whenParsinThenExpectedFormatIsReturned() {
        String result = currencyFormatter.parse(MONEY_INTEGER);

        assertThat(result).isEqualTo(MONEY_FORMATTED);
    }
}