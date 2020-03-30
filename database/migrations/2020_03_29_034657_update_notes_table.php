<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class UpdateNotesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('notes', function (Blueprint $table) {
          $table->timestamp('expiry')->default('2020');
          $table->integer('expiresViews')->default(2);
          $table->integer('viewCount')->default(0);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
      Schema::table('notes', function (Blueprint $table) {
        $table->dropColumn(['viewCount', 'expiresViews', 'expiry']);
      });
    }
}
