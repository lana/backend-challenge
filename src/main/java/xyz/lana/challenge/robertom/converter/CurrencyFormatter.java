package xyz.lana.challenge.robertom.converter;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

import java.text.DecimalFormat;
import java.text.DecimalFormatSymbols;
import java.util.Locale;

@Component
@Slf4j
public class CurrencyFormatter {

    public String parse(double money) {
        double currencyAmount = money / 100;

        DecimalFormatSymbols otherSymbols = new DecimalFormatSymbols(Locale.GERMAN);
        otherSymbols.setDecimalSeparator('.');
        otherSymbols.setGroupingSeparator(',');
        DecimalFormat df = new DecimalFormat("#,##0.00€", otherSymbols);

        return df.format(currencyAmount);
    }
}
