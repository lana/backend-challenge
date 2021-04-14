package xyz.lana.challenge.robertom.api.request;

import lombok.Data;
import xyz.lana.challenge.robertom.model.Item;

@Data
public class AddItemRequest {

    private Item code;
    private int quantity;

}
