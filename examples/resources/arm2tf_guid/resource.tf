# The following example shows how to generate a determinsitic guid for tag on an Azure Storage Account.

resource "arm2tf_guid" "storage" {
  input = [
    azurerm_resource_group.rg.name,
    "production"
  ]
}

resource "azurerm_storage_account" "my_storage_account" {
  name                     = "strgaccount"
  location                 = azurerm_resource_group.rg.location
  resource_group_name      = azurerm_resource_group.rg.name
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    identifier = arm2tf_guid.storage.id
  }
}
