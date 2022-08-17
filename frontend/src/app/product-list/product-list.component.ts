import { Observable } from "rxjs";
import { Product } from "../product";
import { Component, OnInit } from "@angular/core";
import { Router } from '@angular/router';
import { ProductService } from "../product.service";

@Component({
  selector: "app-product-list",
  templateUrl: "./product-list.component.html",
  styleUrls: ["./product-list.component.css"]
})
export class ProductListComponent implements OnInit {
  products: Observable<Product[]>;

  constructor(private productService: ProductService,
    private router: Router) {}

  ngOnInit() {
    this.reloadData();
  }

  reloadData() {
    this.products = this.productService.getProductsList();
  }

  deleteProduct(id: string) {
  
    this.productService.deleteProduct(id)
      .subscribe(
        data => {
          console.log(data);
          this.reloadData();
        },
        error => console.log(error));
  }
  updateProduct(id: string){
    this.router.navigate(['update', id]);
  }
  productDetails(id: string){
    this.router.navigate(['details', id]);
  }
}