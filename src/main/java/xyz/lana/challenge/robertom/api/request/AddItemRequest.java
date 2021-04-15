package xyz.lana.challenge.robertom.api.request;

import lombok.Data;
import xyz.lana.challenge.robertom.model.Item;
import xyz.lana.challenge.robertom.validation.ValueOfEnum;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;

@Data
public class AddItemRequest {

    @NotNull
    @ValueOfEnum(enumClass = Item.class)
    private String code;

    @Min(1)
    private int quantity = 1;

}
