# The following example shows how to generate a unique name for an Azure Storage Account.

resource "arm2tf_unique_string" "storage" {
  input = [azurerm_resource_group.rg.name]
}

resource "azurerm_storage_account" "my_storage_account" {
  name                     = "strg${arm2tf_unique_string.storage.id}"
  location                 = azurerm_resource_group.rg.location
  resource_group_name      = azurerm_resource_group.rg.name
  account_tier             = "Standard"
  account_replication_type = "LRS"
}
